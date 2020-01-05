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
	"github.com/greatgodapollo/Vi/Status"
	"io"
)

type CommandContext struct {
	Session       *discordgo.Session
	Event         *discordgo.MessageCreate
	Manager       *CommandManager
	StatusManager *Status.StatusManager
	Message       *discordgo.Message
	User          *discordgo.User
	Channel       *discordgo.Channel
	Guild         *discordgo.Guild
	Member        *discordgo.Member
}

func (ctx *CommandContext) Reply(message string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSend(ctx.Channel.ID, message)
}

func (ctx *CommandContext) ReplyEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendEmbed(ctx.Channel.ID, embed)
}

func (ctx *CommandContext) ReplyFile(filename string, file io.Reader) (*discordgo.Message, error) {
	return ctx.Session.ChannelFileSend(ctx.Channel.ID, filename, file)
}
