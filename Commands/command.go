package Commands

import "github.com/greatgodapollo/Vi/Shared"

type CommandFunc func(CommandContext, []string) error

type Command struct {
	Name            string
	Description     string
	OwnerOnly       bool
	Hidden          bool
	UserPermissions Shared.Permission
	BotPermissions  Shared.Permission
	Type            CommandType
	Run             CommandFunc
}

type CommandType int

const (
	CommandTypePM CommandType = iota
	CommandTypeGuild
	CommandTypeEverywhere
)
