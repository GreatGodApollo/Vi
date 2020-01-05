package Commands

import (
	"encoding/json"
	"github.com/greatgodapollo/Vi/Shared"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var (
	tags map[string]string
)

func NewTagCommand() *Command {
	return &Command{
		Name:            "tag",
		Description:     "Get a tag",
		OwnerOnly:       false,
		Hidden:          false,
		UserPermissions: 0,
		BotPermissions:  Shared.PermissionMessagesSend,
		Type:            CommandTypeEverywhere,
		Run:             TagCommand,
	}
}

func TagCommand(ctx CommandContext, args []string) error {
	var err error
	if tag, has := tags[args[0]]; has {
		_, err = ctx.Reply(tag)
	} else {
		_, err = ctx.Reply("Tag does not exist")
	}
	return err
}

func LoadTags(f string, log *logrus.Logger) {
	file, err := os.Open(f)
	defer file.Close()

	if err != nil {
		log.Fatal(err.Error())
	}
	byteValues, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValues, &tags)
}
