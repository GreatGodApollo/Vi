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
	"github.com/GreatGodApollo/Vi/Database"
	"github.com/GreatGodApollo/Vi/Shared"
	"strconv"
	"strings"
)

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
	ProcessArgs:     OwnerArgsFunc,
}

// A OwnerCommandsArg is passed into a CommandContext. It provides the necessary information for a help command to run.
type OwnerCommandArgs struct {
	// The name of the command the user is searching for
	Option string

	// The rest of the arguments
	Rest []string
}

// OwnerArgsFunc is a CommandArgFunc
// It returns the proper OwnerCommandArgs struct given the args provided
// It returns an empty struct if no args are provided
func OwnerArgsFunc(args []string) interface{} {
	if len(args) == 1 {
		return OwnerCommandArgs{args[0], nil}
	} else if len(args) > 1 {
		return OwnerCommandArgs{args[0], args[1:]}
	}
	return OwnerCommandArgs{}
}

// OwnerCommandFunc is a CommandRunFunc.
// It contains all of the owner commands
// It returns an error if any occurred.
//
// Usage: {prefix}owner {reloadtags|updateSuggestion} [arguments]
func OwnerCommandFunc(ctx CommandContext, args []string) error {
	argStruct := ctx.Args.(OwnerCommandArgs)
	if argStruct.Option != "" {
		switch argStruct.Option {
		case "reloadtags":
			{
				LoadTags("tags.json", ctx.Manager.Logger)
				_, err := ctx.Reply("Tags reloaded!")
				return err
			}
		case "updateSuggestion", "us":
			{
				switch argStruct.Rest[0] {
				case "status", "s":
					{
						id, err := strconv.Atoi(argStruct.Rest[1])
						if err != nil {
							_, err = ctx.Reply("Invalid suggestion ID")
							return err
						}
						sm := ctx.Manager.DB.First(&Database.Suggestion{}, id)
						var status int
						var statuss string
						var color int
						switch argStruct.Rest[2] {
						case "pending", "p":
							{
								status = Database.SuggestionStatusPending
								statuss = "Pending"
								color = Shared.COLOR
							}
						case "accepted", "a":
							{
								status = Database.SuggestionStatusAccepted
								statuss = "Accepted"
								color = 0x00FF00
							}
						case "denied", "d":
							{
								status = Database.SuggestionStatusDenied
								statuss = "Denied"
								color = 0xFF0000
							}
						default:
							{
								_, err := ctx.Reply("Invalid status code")
								return err
							}
						}
						var suggestion Database.Suggestion
						sm.Find(&suggestion)
						suggestion.Status = status
						sm.Save(suggestion)
						msg, err := ctx.Session.ChannelMessage(suggestion.ChannelId, suggestion.MessageId)
						if err != nil {
							return err
						}

						embedBuilder := Shared.NewEmbed().
							SetTitle(msg.Embeds[0].Title).
							SetAuthor(msg.Embeds[0].Author.Name, msg.Embeds[0].Author.IconURL).
							SetColor(color).
							SetDescription(msg.Embeds[0].Description).
							AddInlineField("Status", statuss).
							SetFooter(msg.Embeds[0].Footer.Text)

						if msg.Embeds[0].Fields[1] != nil {
							embedBuilder.AddInlineField("Message", msg.Embeds[0].Fields[1].Value)
						}

						_, err = ctx.Session.ChannelMessageEditEmbed(suggestion.ChannelId, suggestion.MessageId, embedBuilder.MessageEmbed)
						if err != nil {
							return err
						}
						_, err = ctx.Reply("Status updated!")
						if err != nil {
							return err
						}
						return nil
					}
				case "delete", "d":
					{
						id, err := strconv.Atoi(argStruct.Rest[1])
						if err != nil {
							_, err = ctx.Reply("Invalid suggestion ID")
							return err
						}
						sm := ctx.Manager.DB.First(&Database.Suggestion{}, id)
						var suggestion Database.Suggestion
						sm.Find(&suggestion)

						if (suggestion != Database.Suggestion{}) {
							err = ctx.Session.ChannelMessageDelete(suggestion.ChannelId, suggestion.MessageId)
							if err != nil {
								return err
							}

							ctx.Manager.DB.Where("id=?", id).Delete(&Database.Suggestion{})
							_, err = ctx.Reply("Suggestion deleted!")
							return err
						} else {
							_, err = ctx.Reply("Invalid suggestion ID")
							return err
						}
					}
				case "message", "m":
					{
						id, err := strconv.Atoi(argStruct.Rest[1])
						if err != nil {
							_, err = ctx.Reply("Invalid suggestion ID")
							return err
						}
						msgs := strings.Join(argStruct.Rest[2:], " ")
						sm := ctx.Manager.DB.First(&Database.Suggestion{}, id)
						var suggestion Database.Suggestion
						sm.Find(&suggestion)
						suggestion.Message = msgs
						sm.Save(&suggestion)

						msg, err := ctx.Session.ChannelMessage(suggestion.ChannelId, suggestion.MessageId)
						if err != nil {
							return err
						}

						embedBuilder := Shared.NewEmbed().
							SetTitle(msg.Embeds[0].Title).
							SetAuthor(msg.Embeds[0].Author.Name, msg.Embeds[0].Author.IconURL).
							SetColor(msg.Embeds[0].Color).
							SetDescription(msg.Embeds[0].Description).
							AddInlineField("Status", msg.Embeds[0].Fields[0].Value).
							AddInlineField("Message", msgs).
							SetFooter(msg.Embeds[0].Footer.Text)

						_, err = ctx.Session.ChannelMessageEditEmbed(suggestion.ChannelId, suggestion.MessageId, embedBuilder.MessageEmbed)
						if err != nil {
							return err
						}
						_, err = ctx.Reply("Message updated!")
						if err != nil {
							return err
						}
						return nil
					}
				}
			}
		}
	}
	_, err := ctx.Reply("To be implemented!")
	return err
}
