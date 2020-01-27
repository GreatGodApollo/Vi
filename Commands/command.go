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

// A CommandFunc is ran whenever a CommandManager gets a message supposed to run the given command.
type CommandFunc func(CommandContext, []string) error

// A Command represents any given command contained in a bot.
type Command struct {
	// The name of the command (What it will be triggered by).
	Name string

	// Command aliases
	Aliases []string

	// The command's description.
	Description string

	// If the command is only able to be ran by an owner.
	OwnerOnly bool

	// If the command is hidden from help.
	Hidden bool

	// The permissions the user is required to have to execute the command.
	UserPermissions Shared.Permission

	// The permissions the bot is required to have to execute the command.
	BotPermissions Shared.Permission

	// The CommandType designates where the command can be ran.
	Type CommandType

	// The function that will be executed whenever a message fits the criteria to execute the command.
	Run CommandFunc
}

// A CommandType represents the locations commands can be used.
type CommandType int

const (
	// A Command that is only supposed to run in a personal message
	CommandTypePM CommandType = iota

	// A command that is only supposed to run in a Guild
	CommandTypeGuild

	// A Command that is able to run anywhere
	CommandTypeEverywhere
)
