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
	"github.com/GreatGodApollo/Vi/Shared"
)

// NewInviteCommand returns a new InviteCommand for use in a CommandManager.
// It returns a Command struct.
func NewInviteCommand() *Command {
	return &Command{
		Name:            "invite",
		Description:     "Invite Me!",
		OwnerOnly:       false,
		Hidden:          false,
		UserPermissions: 0,
		BotPermissions:  Shared.PermissionMessagesSend | Shared.PermissionMessagesEmbedLinks,
		Type:            CommandTypeEverywhere,
		Run:             InviteCommand,
	}
}

// InviteCommand is a CommandRunFunc.
// It supplies the user an invite to the bot.
// It returns an error if any occurred.
func InviteCommand(ctx CommandContext, args []string) error {
	embed := Shared.NewEmbed().
		SetTitle("Invite Me!").
		SetColor(Shared.COLOR).
		SetDescription("Vi is a multi-functional Discord bot written in GoLang using the DiscordGo library."+
			"This bot was authored by `apollo#9292` and is available at https://l.brettb.xyz/vi").
		AddField("Invite URL", fmt.Sprintf("https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot&permissions=379976", ctx.Session.State.User.ID))
	_, err := ctx.ReplyEmbed(embed.MessageEmbed)
	return err
}
