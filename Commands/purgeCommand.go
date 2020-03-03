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
	"strconv"
	"time"
)

var PurgeCommand = &Command{
	Name:            "purge",
	Aliases:         []string{"rm"},
	Description:     "Delete a bunch of messages",
	OwnerOnly:       false,
	Hidden:          false,
	UserPermissions: Shared.PermissionMessagesManage,
	BotPermissions:  Shared.PermissionMessagesSend | Shared.PermissionMessagesManage,
	Type:            CommandTypeEverywhere,
	Run:             PurgeCommandFunc,
	ProcessArgs:     PurgeArgsFunc,
}

// A PurgeCommandArgs is passed into a CommandContext. It provides the necessary information for a purge command to run.
type PurgeCommandArgs struct {
	Number int
	Rest   []string
}

// PurgeArgsFunc is a CommandArgFunc
// It returns the proper PurgeCommandArgs struct given the args provided
// It returns an empty struct if no args are provided
func PurgeArgsFunc(args []string) interface{} {
	if len(args) == 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return PurgeCommandArgs{Number: 0, Rest: nil}
		} else {
			return PurgeCommandArgs{Number: i, Rest: nil}
		}
	} else if len(args) > 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return PurgeCommandArgs{Number: 0, Rest: args[1:]}
		} else {
			return PurgeCommandArgs{Number: i, Rest: args[1:]}
		}
	} else {
		return PurgeCommandArgs{}
	}
}

// PurgeCommandFunc is a CommandRunFunc.
// It deletes 'x' number of messages from a channel.
// It returns an error if any occurred.
//
// Usage: {prefix}purge {num}
func PurgeCommandFunc(ctx CommandContext, args []string) error {
	argStruct := ctx.Args.(PurgeCommandArgs)
	err := ctx.PurgeMessages(argStruct.Number)
	if err != nil {
		if err.Error() == "too many messages" {
			_, e := ctx.Reply(":x: You can delete a max of 100 messages :x:")
			return e
		} else if err.Error() == "must supply a number" {
			_, e := ctx.Reply(":x: You must specify a number of messages to delete :x:")
			return e
		} else {
			return err
		}
	}
	embedBuilder := Shared.NewEmbed().
		SetFooter("Initiated by: " + ctx.User.Username + "#" + ctx.User.Discriminator).
		SetTitle("Purged " + strconv.Itoa(argStruct.Number) + " messages!").
		SetColor(Shared.COLOR)
	m, err := ctx.ReplyEmbed(embedBuilder.MessageEmbed)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Second)
	err = ctx.Session.ChannelMessageDelete(m.ChannelID, m.ID)
	return err
}
