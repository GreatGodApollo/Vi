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

import "github.com/GreatGodApollo/Vi/Shared"

var OwnerCommand = &Command{
	Name:            "owner",
	Aliases:         []string{"o"},
	Description:     "The general owner command",
	OwnerOnly:       true,
	Hidden:          false,
	UserPermissions: 0,
	BotPermissions:  Shared.PermissionMessagesSend | Shared.PermissionMessagesEmbedLinks,
	Type:            CommandTypeEverywhere,
	Run:             OwnerCommandFunc,
}

// OwnerCommandFunc is a CommandRunFunc.
// It currently has no use.
// It returns an error if any occurred.
func OwnerCommandFunc(ctx CommandContext, args []string) error {
	if len(args) > 0 {
		switch args[0] {
		case "reloadtags":
			{
				LoadTags("tags.json", ctx.Manager.Logger)
				_, err := ctx.Reply("Tags reloaded!")
				return err
			}
		}
	}
	_, err := ctx.Reply("To be implemented!")
	return err
}
