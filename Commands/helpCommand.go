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
