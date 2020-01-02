package Commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Shared"
)

func NewAboutCommand() *Command {
	return &Command{
		Name:            "about",
		Description:     "Get some information about the bot",
		OwnerOnly:       false,
		Hidden:          false,
		UserPermissions: 0,
		BotPermissions:  Shared.PermissionMessagesSend + Shared.PermissionMessagesEmbedLinks,
		Type:            CommandTypeEverywhere,
		Run:             AboutCommand,
	}
}

func AboutCommand(ctx CommandContext, args []string) error {
	embed := &discordgo.MessageEmbed{
		Title: "About Vi",
		Color: 0,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Created by apollo#9292",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Language",
				Value:  "GoLang",
				Inline: true,
			},
			{
				Name:   "Github",
				Value:  "https://github.com/greatgodapollo/vi",
				Inline: true,
			},
		},
	}

	_, err := ctx.ReplyEmbed(embed)
	return err
}
