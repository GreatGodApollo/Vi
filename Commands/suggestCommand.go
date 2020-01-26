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
	"github.com/GreatGodApollo/Vi/Shared"
	"strings"
)

func NewSuggestCommand() *Command {
	return &Command{
		Name:            "suggest",
		Description:     "Suggest a thing for the bot!",
		OwnerOnly:       false,
		Hidden:          false,
		UserPermissions: 0,
		BotPermissions:  Shared.PermissionMessagesSend,
		Type:            CommandTypeEverywhere,
		Run:             SuggestCommand,
	}
}

func SuggestCommand(ctx CommandContext, args []string) error {
	if ctx.Manager.Config.Miscellaneous.SuggestionChannel == "" {
		_, err := ctx.Reply(":x: Suggesting is not enabled! :x:")
		return err
	}
	if len(args) <= 1 {
		_, err := ctx.Reply(":x: You need to type something actually worth suggesting! :x:")
		return err
	}
	embedBuilder := Shared.NewEmbed()
	embedBuilder.SetColor(Shared.COLOR)
	embedBuilder.SetAuthor("Suggestion from: "+ctx.User.Username+"#"+ctx.User.Discriminator, ctx.User.AvatarURL("1024"))
	embedBuilder.SetDescription(strings.Join(args, " "))
	_, err := ctx.Session.ChannelMessageSendEmbed(ctx.Manager.Config.Miscellaneous.SuggestionChannel, embedBuilder.MessageEmbed)

	if err != nil {
		return err
	}
	_, err = ctx.Reply(":white_check_mark: Suggestion sent!")
	return err
}
