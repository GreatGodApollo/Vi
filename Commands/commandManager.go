/*
 * Vi - A Discord Bot written in Go
 * Copyright (C) 2019  Brett Bender
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// The Commands package both contains the CommandManager framework and the bot commands.
// Everything is pretty modular and can be adapted to your own use cases.
package Commands

import (
	"github.com/GreatGodApollo/Vi/Configuration"
	"github.com/GreatGodApollo/Vi/Shared"
	"github.com/GreatGodApollo/Vi/Status"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"strings"
)

// CommandHandler works as the CommandManager's message listener.
// It returns nothing.
func (cmdm *CommandManager) CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Bot && cmdm.IgnoreBots {
		return
	}

	var prefix string
	var contains bool
	var err error
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

	if command, exist, _ := cmdm.GetCommand(cmd[0]); exist {
		var inDm bool
		if channel.Type == discordgo.ChannelTypeDM {
			inDm = true
		}

		// Check UserPermissions
		if command.Type != CommandTypePM && !inDm && !Shared.CheckPermissions(s, m.GuildID, m.Author.ID, command.UserPermissions) {
			if Shared.CheckPermissions(s, m.GuildID, s.State.User.ID, Shared.PermissionMessagesEmbedLinks) {
				embed := &discordgo.MessageEmbed{
					Title:       "Insufficient Permissions!",
					Description: "You don't have the required permissions to run this command!",
					Color:       0xff0000,
				}

				if !command.Hidden {
					_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				}
			} else {
				if !command.Hidden {
					_, err = s.ChannelMessageSend(m.ChannelID, ":x: You don't have the correct permissions to run this command! :x:")
				}
			}
			if err != nil {
				cmdm.OnErrorFunc(cmdm, err)
			}
			cmdm.Logger.Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		// Check BotPermissions
		if command.Type != CommandTypePM && !inDm && !Shared.CheckPermissions(s, m.GuildID, s.State.User.ID, command.BotPermissions) {
			if Shared.CheckPermissions(s, m.GuildID, s.State.User.ID, Shared.PermissionMessagesEmbedLinks) {
				embed := &discordgo.MessageEmbed{
					Title:       "Insufficient Permissions!",
					Description: "I don't have the correct permissions to run this command!",
					Color:       0xff0000,
				}

				if !command.Hidden {
					_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
				}
			} else {
				if !command.Hidden {
					_, err = s.ChannelMessageSend(m.ChannelID, ":x: I don't have the correct permissions to run this command! :x:")
				}
			}

			if err != nil {
				cmdm.OnErrorFunc(cmdm, err)
			}
			cmdm.Logger.Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		if channel.Type == discordgo.ChannelTypeDM && command.Type == CommandTypeGuild {
			embed := &discordgo.MessageEmbed{
				Title:       "Invalid Channel!",
				Description: "You cannot run this command in a private message.",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			if err != nil {
				cmdm.OnErrorFunc(cmdm, err)
			}
			cmdm.Logger.Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		} else if channel.Type == discordgo.ChannelTypeGuildText && command.Type == CommandTypePM {
			embed := &discordgo.MessageEmbed{
				Title:       "Invalid Channel!",
				Description: "You cannot run this command in a guild.",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			if err != nil {
				cmdm.OnErrorFunc(cmdm, err)
			}
			cmdm.Logger.Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		if command.OwnerOnly && !cmdm.IsOwner(m.Author.ID) {
			embed := &discordgo.MessageEmbed{
				Title:       "You can't run that command!",
				Description: "Sorry, only bot owners can run that command",
				Color:       0xff0000,
			}

			if !command.Hidden {
				_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
			}

			if err != nil {
				cmdm.OnErrorFunc(cmdm, err)
			}
			cmdm.Logger.Debugf("P: FALSE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
			return
		}

		cmdm.Logger.Debugf("P: TRUE C: %s[%s] U: %s#%s[%s] M: %s", channel.Name, m.ChannelID, m.Author.Username, m.Author.Discriminator, m.Author.ID, m.Content)
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

		err = command.Run(ctx, cmd[1:])
		if err != nil {
			cmdm.OnErrorFunc(cmdm, err)
		}
	}
}

// AddPrefix adds a new prefix to the CommandManager's prefix list.
// It returns nothing.
func (cmdm *CommandManager) AddPrefix(prefix string) {
	cmdm.Prefixes = append(cmdm.Prefixes, prefix)
}

// RemovePrefix removes a prefix from the CommandManager's prefix list.
// It returns nothing.
func (cmdm *CommandManager) RemovePrefix(prefix string) {
	for i, v := range cmdm.Prefixes {
		if v == prefix {
			cmdm.Prefixes = append(cmdm.Prefixes[:i], cmdm.Prefixes[i+1:]...)
			break
		}
	}
}

// SetPrefixes sets the CommandManager's prefix list.
// It returns nothing.
func (cmdm *CommandManager) SetPrefixes(prefixes []string) {
	cmdm.Prefixes = prefixes
}

// GetPrefixes gets the CommandManager's prefix list.
// It returns a string array.
func (cmdm *CommandManager) GetPrefixes() []string {
	return cmdm.Prefixes
}

// AddNewCommand adds a new command to the CommandManager's command list.
// It returns nothing.
func (cmdm *CommandManager) AddNewCommand(name string, aliases []string, desc string, owneronly, hidden bool, userperms, botperms Shared.Permission,
	cmdType CommandType, run CommandFunc) {
	var cmd *Command
	if _, exists, _ := cmdm.GetCommand(name); !exists {
		cmd = &Command{
			name, aliases, desc, owneronly, hidden, userperms, botperms, cmdType, run,
		}
	}
	*cmdm.Commands = append(*cmdm.Commands, cmd)
}

// AddCommand adds an existent command to the CommandManager's command list.
// It returns nothing.
func (cmdm *CommandManager) AddCommand(cmd *Command) {
	if _, exists, _ := cmdm.GetCommand(cmd.Name); !exists {
		*cmdm.Commands = append(*cmdm.Commands, cmd)
	}
}

func (cmdm *CommandManager) GetCommand(name string) (cmd *Command, exists bool, index int) {
	for i, c := range *cmdm.Commands {
		if c.Name == name {
			return c, true, i
		}
		for _, a := range c.Aliases {
			if a == name {
				return c, true, i
			}
		}
	}
	return nil, false, 0
}

// RemoveCommand removes a command from the CommandManager's command list.
// It returns nothing.
func (cmdm *CommandManager) RemoveCommand(name string) {
	if _, exists, index := cmdm.GetCommand(name); exists {
		*cmdm.Commands = RemoveCommandFromSlice(*cmdm.Commands, index)
	}
}

// IsOwner checks if a user ID is is in the owner list.
// It returns a bool.
func (cmdm *CommandManager) IsOwner(id string) bool {
	for _, o := range cmdm.Owners {
		if id == o {
			return true
		}
	}
	return false
}

// NewCommandManager instantiates a new CommandManager.
// It returns a CommandManager.
func NewCommandManager(c Configuration.Configuration, sm *Status.StatusManager, l *logrus.Logger, ignoreBots bool, errorFunc CommandManagerOnErrorFunc) CommandManager {
	return CommandManager{
		Config:        c,
		Prefixes:      c.Bot.Prefixes,
		Owners:        c.Bot.Owners,
		StatusManager: sm,
		Logger:        l,
		Commands:      &[]*Command{},
		IgnoreBots:    ignoreBots,
		OnErrorFunc:   errorFunc,
	}
}

// A CommandManager represents a set of prefixes, owners and commands, with some extra utility to create a command handler.
type CommandManager struct {
	// The bot configuration
	Config Configuration.Configuration

	// The array of prefixes a CommandManager will respond to.
	Prefixes []string

	// The array of IDs that will be considered a bot owner.
	Owners []string

	// The bot's StatusManager.
	StatusManager *Status.StatusManager

	// The bot instance Logger.
	Logger *logrus.Logger

	// The map of Commands in the CommandManager.
	Commands *[]*Command

	// If the CommandManager ignores bots or not.
	IgnoreBots bool

	// The function that will be ran when the CommandManager encounters an error.
	OnErrorFunc CommandManagerOnErrorFunc
}

// A CommandManagerOnErrorFunc is a function that will run whenever the CommandManager encounters an error.
type CommandManagerOnErrorFunc func(cmdm *CommandManager, err error)

func RemoveCommandFromSlice(s []*Command, i int) []*Command {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
