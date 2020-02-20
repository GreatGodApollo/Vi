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
	"github.com/GreatGodApollo/Vi/Status"
	"github.com/bwmarrin/discordgo"
	"io"
)

// A CommandContext is passed to a CommandRunFunc. It contains the information needed for a command to execute.
type CommandContext struct {
	// The connection to Discord.
	Session *discordgo.Session

	// The event that fired the CommandHandler.
	Event *discordgo.MessageCreate

	// The CommandManager that handled this command.
	Manager *CommandManager

	// The custom args struct for this command
	Args interface{}

	// The bot's StatusManager.
	StatusManager *Status.StatusManager

	// The Message that fired this event.
	Message *discordgo.Message

	// The User that fired this event.
	User *discordgo.User

	// The Channel the event was fired in.
	Channel *discordgo.Channel

	// The guild the Channel belongs to.
	Guild *discordgo.Guild

	// The User's guild member.
	Member *discordgo.Member
}

// Reply sends a message to the channel a CommandContext was initiated for.
func (ctx *CommandContext) Reply(message string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSend(ctx.Channel.ID, message)
}

// ReplyEmbed sends an embed to the channel a CommandContext was initiated for.
func (ctx *CommandContext) ReplyEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendEmbed(ctx.Channel.ID, embed)
}

// ReplyFile sends a file to the channel a CommandContext was initiated for.
func (ctx *CommandContext) ReplyFile(filename string, file io.Reader) (*discordgo.Message, error) {
	return ctx.Session.ChannelFileSend(ctx.Channel.ID, filename, file)
}
