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
	"time"
)

var PingCommand = &Command{
	Name:            "ping",
	Description:     "Check if the bot is alive",
	OwnerOnly:       false,
	Hidden:          false,
	UserPermissions: 0,
	BotPermissions:  Shared.PermissionMessagesSend,
	Type:            CommandTypeEverywhere,
	Run:             PingCommandFunc,
}

// PingCommandFunc is a CommandRunFunc.
// It supplies the user a message if the bot is alive.
// It returns an error if any occurred.
//
// Usage: {prefix}ping
func PingCommandFunc(ctx CommandContext, args []string) error {
	initTime := time.Now()
	m, err := ctx.Reply("Pinging...")
	if err != nil {
		return err
	}
	_, err = Shared.EditMessageText(ctx.Session, m, "Pong! Took: "+time.Since(initTime).String())
	return err
}
