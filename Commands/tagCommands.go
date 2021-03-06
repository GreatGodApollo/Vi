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
	"fmt"
	"github.com/GreatGodApollo/Vi/Shared"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var (
	tags map[string]string
)

var TagCommand = &Command{
	Name:            "tag",
	Aliases:         []string{"t"},
	Description:     "Get a tag",
	OwnerOnly:       false,
	Hidden:          false,
	UserPermissions: 0,
	BotPermissions:  Shared.PermissionMessagesSend,
	Type:            CommandTypeEverywhere,
	Run:             TagCommandFunc,
	ProcessArgs:     TagArgsFunc,
}

var TagsCommand = &Command{
	Name:            "tags",
	Aliases:         []string{"ts"},
	Description:     "Get a list of tags",
	OwnerOnly:       false,
	Hidden:          false,
	UserPermissions: 0,
	BotPermissions:  Shared.PermissionMessagesSend | Shared.PermissionMessagesEmbedLinks,
	Type:            CommandTypeEverywhere,
	Run:             TagsCommandFunc,
}

// A TagCommandArgs is passed into a CommandContext. It provides the necessary information for a tag command to run.
type TagCommandArgs struct {
	// The name of the tag the user is searching for
	Tag string
}

// TagArgsFunc is a CommandArgFunc
// It returns the proper TagCommandArgs struct given the args provided
// It returns an empty struct if no args are provided
func TagArgsFunc(args []string) interface{} {
	if len(args) == 1 {
		return TagCommandArgs{args[0]}
	}
	return TagCommandArgs{}
}

// TagCommandFunc is a CommandRunFunc.
// It supplies the user with the tag description if the tag supplied exists.
// It returns an error if any occurred.
//
// Usage: {prefix}tag <tag>
func TagCommandFunc(ctx CommandContext, args []string) error {
	argStruct := ctx.Args.(TagCommandArgs)
	if argStruct.Tag != "" {
		var err error
		if tag, has := tags[argStruct.Tag]; has {
			_, err = ctx.Reply(tag)
		} else {
			_, err = ctx.Reply(":x: Tag does not exist :x:")
		}
		return err
	} else {
		_, err := ctx.Reply(":x: Please supply a tag :x:")
		return err
	}
}

// TagsCommandFunc is a CommandRunFunc.
// It supplies the user with a list of the initialized tags.
// It returns an error if any occurred.
//
// Usage: {prefix}tags
func TagsCommandFunc(ctx CommandContext, _ []string) error {
	if len(tags) > 0 {
		var list string
		for k, _ := range tags {
			list += fmt.Sprintf("**`%s`**\n", k)
		}
		embed := Shared.NewEmbed().
			SetTitle("Tags:").
			SetDescription(list).
			SetColor(Shared.COLOR)

		_, err := ctx.ReplyEmbed(embed.MessageEmbed)
		return err
	}

	_, err := ctx.Reply(":x: No tags are initialized :x:")
	return err
}

// LoadTags loads the tags from a given file.
// It returns nothing.
func LoadTags(f string, log *logrus.Logger) {
	tags = nil
	file, err := os.Open(f)
	defer file.Close()

	if err != nil {
		log.Fatal(err.Error())
	}
	byteValues, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValues, &tags)
}
