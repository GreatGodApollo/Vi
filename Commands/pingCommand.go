package Commands

import "github.com/greatgodapollo/Vi/Shared"

func NewPingCommand() *Command {
	return &Command{
		Name:            "ping",
		Description:     "Check if the bot is alive",
		OwnerOnly:       false,
		Hidden:          false,
		UserPermissions: 0,
		BotPermissions:  Shared.PermissionMessagesSend,
		Type:            CommandTypeEverywhere,
		Run:             PingCommand,
	}
}

func PingCommand(ctx CommandContext, args []string) error {
	_, err := ctx.Reply("Pong!")
	return err
}
