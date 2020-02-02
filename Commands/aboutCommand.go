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
)

var AboutCommand = &Command{
	Name:            "about",
	Description:     "Get some information about the bot",
	OwnerOnly:       false,
	Hidden:          false,
	UserPermissions: 0,
	BotPermissions:  Shared.PermissionMessagesSend | Shared.PermissionMessagesEmbedLinks,
	Type:            CommandTypeEverywhere,
	Run:             AboutCommandFunc,
}

// AboutCommandFunc is a CommandRunFunc.
// It supplies the user who runs it information about the bot.
// It returns an error if any occurred.
//
// Usage: {prefix}about
func AboutCommandFunc(ctx CommandContext, args []string) error {
	e := Shared.NewEmbed().
		SetTitle("About Vi").
		SetColor(Shared.COLOR).
		SetFooter("Created by apollo#9292").
		AddInlineField("Language", "GoLang").
		AddInlineField("Github", "https://l.brettb.xyz/vi").
		AddInlineField("Version", Shared.VERSION)

	_, err := ctx.ReplyEmbed(e.MessageEmbed)
	return err
}
