package Commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Shared"
	"sort"
	"strings"
)

func NewHelpCommand() *Command {
	return &Command{
		Name:            "help",
		Description:     "Get some help with the bot.",
		OwnerOnly:       false,
		Hidden:          false,
		UserPermissions: 0,
		BotPermissions:  Shared.PermissionMessagesSend + Shared.PermissionMessagesEmbedLinks,
		Type:            CommandTypeEverywhere,
		Run:             HelpCommand,
	}
}

func HelpCommand(ctx CommandContext, args []string) error {

	if len(args) > 0 {
		if command, has := ctx.Manager.Commands[strings.ToLower(args[0])]; has {
			if command.Hidden {
				return nil
			}

			var (
				ownerOnlyString string
				typeString      string
			)

			if command.OwnerOnly {
				ownerOnlyString = "Yes"
			} else {
				ownerOnlyString = "No"
			}

			switch command.Type {
			case CommandTypePM:
				{
					typeString = "Private"
				}
			case CommandTypeGuild:
				{
					typeString = "Guild-only"
				}
			case CommandTypeEverywhere:
				{
					typeString = "Anywhere"
				}
			}

			embed := &discordgo.MessageEmbed{
				Title: fmt.Sprintf("Help for `%s`!", args[0]),
				Color: Shared.COLOR,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Description:",
						Value: command.Description,
					},
					{
						Name:   "Owner Only?",
						Value:  ownerOnlyString,
						Inline: true,
					},
					{
						Name:   "Usage?",
						Value:  typeString,
						Inline: true,
					},
				},
			}

			_, err := ctx.ReplyEmbed(embed)
			return err
		}
	}
	m := ctx.Manager.Commands

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var list string
	for _, k := range keys {
		cmd := m[k]
		if !cmd.Hidden {
			list += fmt.Sprintf("**%s** - `%s`\n", cmd.Name, cmd.Description)
		}
	}

	var footer strings.Builder

	if len(ctx.Manager.Commands) == 1 {
		footer.WriteString("There is 1 command.")
	} else {
		footer.WriteString(fmt.Sprintf("There are %d commands.", len(ctx.Manager.Commands)))
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Commands:",
		Description: list,
		Color:       Shared.COLOR,
		Footer: &discordgo.MessageEmbedFooter{
			Text: footer.String(),
		},
	}

	_, err := ctx.ReplyEmbed(embed)
	return err
}
