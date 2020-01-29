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
		case "updateSuggestion", "us":
			{
				switch args[1] {
				case "status", "s":
					{
						id, err := strconv.Atoi(args[2])
						if err != nil {
							ctx.Reply("Invalid suggestion ID")
						}
						sm := ctx.Manager.DB.First(&Database.Suggestion{}, id)
						var status int
						var statuss string
						var color int
						switch args[3] {
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
								if err != nil {
									return err
								}
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
						id, err := strconv.Atoi(args[2])
						if err != nil {
							ctx.Reply("Invalid suggestion ID")
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
				}
			}
		}
	}
	_, err := ctx.Reply("To be implemented!")
	return err
}
