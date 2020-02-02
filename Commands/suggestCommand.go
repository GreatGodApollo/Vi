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
	"fmt"
	"github.com/GreatGodApollo/Vi/Database"
	"github.com/GreatGodApollo/Vi/Shared"
	"strings"
)

var SuggestCommand = &Command{
	Name:            "suggest",
	Description:     "Suggest a thing for the bot!",
	OwnerOnly:       false,
	Hidden:          false,
	UserPermissions: 0,
	BotPermissions:  Shared.PermissionMessagesSend,
	Type:            CommandTypeEverywhere,
	Run:             SuggestCommandFunc,
}

// SuggestCommandFunc is a CommandRunFunc.
// It submits a suggestion to the channel specific in the config.
// It returns an error if any occurred.
//
// Usage: {prefix}suggest <suggestion>
func SuggestCommandFunc(ctx CommandContext, args []string) error {
	if ctx.Manager.Config.Miscellaneous.SuggestionChannel == "" {
		_, err := ctx.Reply(":x: Suggesting is not enabled! :x:")
		return err
	}
	if len(args) <= 1 {
		_, err := ctx.Reply(":x: You need to type something actually worth suggesting! :x:")
		return err
	}
	s := strings.Join(args, " ")
	if len(s) > 512 {
		_, err := ctx.Reply(":x: Your suggestion needs to be less than 512 characters! :x:")
		return err
	}
	suggestion := &Database.Suggestion{
		MessageId:  "",
		ChannelId:  "",
		Suggestion: s,
		Status:     Database.SuggestionStatusPending,
		Message:    "",
	}
	ctx.Manager.DB.Create(suggestion)
	embedBuilder := Shared.NewEmbed()
	embedBuilder.SetColor(Shared.COLOR)
	embedBuilder.SetAuthor("Suggestion from: "+ctx.User.Username+"#"+ctx.User.Discriminator, ctx.User.AvatarURL("1024"))
	embedBuilder.SetDescription(s)
	embedBuilder.AddInlineField("Status", "Pending")
	embedBuilder.AddInlineField("Message", "nil")
	embedBuilder.SetFooter(fmt.Sprintf("Suggestion ID: %v | User ID: %v", suggestion.ID, ctx.User.ID))
	m, err := ctx.Session.ChannelMessageSendEmbed(ctx.Manager.Config.Miscellaneous.SuggestionChannel, embedBuilder.MessageEmbed)
	if err != nil {
		return err
	}
	suggestion.ChannelId = m.ChannelID
	suggestion.MessageId = m.ID
	ctx.Manager.DB.Save(suggestion)
	_, err = ctx.Reply(":white_check_mark: Suggestion sent!")
	return err
}
