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
		Color: Shared.COLOR,
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
			{
				Name:   "Version",
				Value:  Shared.VERSION,
				Inline: true,
			},
		},
	}

	_, err := ctx.ReplyEmbed(embed)
	return err
}
