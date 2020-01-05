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
