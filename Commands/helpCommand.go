package Commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Shared"
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

	var list string
	for _, cmd := range ctx.Manager.Commands {
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
		Footer: &discordgo.MessageEmbedFooter{
			Text: footer.String(),
		},
	}

	_, err := ctx.ReplyEmbed(embed)
	return err
}
