package Commands

import "github.com/greatgodapollo/Vi/Shared"

func NewOwnerCommand() *Command {
	return &Command{
		Name:            "owner",
		Description:     "The general owner command",
		OwnerOnly:       true,
		Hidden:          false,
		UserPermissions: 0,
		BotPermissions:  Shared.PermissionMessagesSend + Shared.PermissionMessagesEmbedLinks,
		Type:            CommandTypeEverywhere,
		Run:             OwnerCommand,
	}
}

func OwnerCommand(ctx CommandContext, args []string) error {
	_, err := ctx.Reply("To be implemented!")
	return err
}
