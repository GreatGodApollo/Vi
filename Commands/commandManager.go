package Commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Configuration"
	"github.com/greatgodapollo/Vi/Shared"
	"github.com/greatgodapollo/Vi/Status"
	"strings"
)

func (cmdm *CommandManager) CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Bot && cmdm.IgnoreBots {
		return
	}

	var prefix string
	var contains bool
	for i := 0; i < len(cmdm.Prefixes); i++ {
		prefix = cmdm.Prefixes[i]
		if strings.HasPrefix(m.Content, prefix) {
			contains = true
			break
		}
	}

	if !contains {
		return
	}

	cmd := strings.Split(strings.TrimPrefix(m.Content, prefix), " ")

	channel, _ := s.Channel(m.ChannelID)

	if command, exist := cmdm.Commands[cmd[0]]; exist {
		var inDm bool
		if channel.Type == discordgo.ChannelTypeDM {
			inDm = true
		}

		// Check UserPermissions
		if command.Type != CommandTypePM && !inDm && !Shared.CheckPermissions(s, m.GuildID, m.Author.ID,
			command.UserPermissions) {
			embed := &discordgo.MessageEmbed{
				Title:       "Insufficient Permissions!",
				Description: "You don't have the required permissions to run this command!",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			return
		}

		// Check BotPermissions
		if command.Type != CommandTypePM && !inDm && !Shared.CheckPermissions(s, m.GuildID, s.State.User.ID,
			command.BotPermissions) {
			embed := &discordgo.MessageEmbed{
				Title:       "Insufficient Permissions!",
				Description: "I don't have the correct permissions to run this command!",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			return
		}

		if channel.Type == discordgo.ChannelTypeDM && command.Type == CommandTypeGuild {
			embed := &discordgo.MessageEmbed{
				Title:       "Invalid Channel!",
				Description: "You cannot run this command in a private message.",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			return
		} else if channel.Type == discordgo.ChannelTypeGuildText && command.Type == CommandTypePM {
			embed := &discordgo.MessageEmbed{
				Title:       "Invalid Channel!",
				Description: "You cannot run this command in a guild.",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			return
		}

		if command.OwnerOnly && !cmdm.IsOwner(m.Author.ID) {
			embed := &discordgo.MessageEmbed{
				Title:       "You can't run that command!",
				Description: "Sorry, only bot owners can run that command",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			return
		}

		guild, _ := s.Guild(m.GuildID)
		member, _ := s.State.Member(m.GuildID, m.Author.ID)

		ctx := CommandContext{
			Session:       s,
			Event:         m,
			Manager:       cmdm,
			StatusManager: cmdm.StatusManager,
			Message:       m.Message,
			User:          m.Author,
			Channel:       channel,
			Guild:         guild,
			Member:        member,
		}

		_ = command.Run(ctx, cmd[1:])
	}
}

func (cmdm *CommandManager) AddPrefix(prefix string) {
	cmdm.Prefixes = append(cmdm.Prefixes, prefix)
}

func (cmdm *CommandManager) RemovePrefix(prefix string) {
	for i, v := range cmdm.Prefixes {
		if v == prefix {
			cmdm.Prefixes = append(cmdm.Prefixes[:i], cmdm.Prefixes[i+1:]...)
			break
		}
	}
}

func (cmdm *CommandManager) SetPrefixes(prefixes []string) {
	cmdm.Prefixes = prefixes
}

func (cmdm *CommandManager) GetPrefixes() []string {
	return cmdm.Prefixes
}

func (cmdm *CommandManager) AddNewCommand(name, desc string, owneronly, hidden bool, userperms, botperms Shared.Permission,
	cmdType CommandType, run CommandFunc) {
	cmdm.Commands[name] = &Command{
		name, desc, owneronly, hidden, userperms, botperms, cmdType, run,
	}
}

func (cmdm *CommandManager) AddCommand(cmd *Command) {
	cmdm.Commands[cmd.Name] = cmd
}

func (cmdm *CommandManager) RemoveCommand(name string) {
	if _, has := cmdm.Commands[name]; has {
		delete(cmdm.Commands, name)
	}
	return
}

func (cmdm *CommandManager) IsOwner(id string) bool {
	for _, o := range cmdm.Owners {
		if id == o {
			return true
		}
	}
	return false
}

func NewCommandManager(c Configuration.Configuration, sm *Status.StatusManager, ignoreBots bool) CommandManager {
	return CommandManager{
		Prefixes:      c.Bot.Prefixes,
		Owners:        c.Bot.Owners,
		StatusManager: sm,
		Commands:      make(map[string]*Command),
		IgnoreBots:    ignoreBots,
	}
}

type CommandManager struct {
	Prefixes      []string
	Owners        []string
	StatusManager *Status.StatusManager
	Commands      map[string]*Command
	IgnoreBots    bool
}
