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
	"github.com/GreatGodApollo/Vi/Shared"
	"github.com/bwmarrin/discordgo"
	"math"
	"strconv"
	"strings"
	"time"
)

var tformat = "Mon, 02 Jan 2006\n15:04:05 MST"

var UserInfoCommand = &Command{
	Name:            "userinfo",
	Aliases:         []string{"ui"},
	Description:     "Returns information about a given user",
	OwnerOnly:       false,
	Hidden:          false,
	UserPermissions: 0,
	BotPermissions:  Shared.PermissionMessagesSend,
	Type:            CommandTypeEverywhere,
	Run:             UserInfoCommandFunc,
}

// UserInfoCommand is a CommandRunFunc.
// It supplies the user information about another user.
// It returns an error if any occurred.
func UserInfoCommandFunc(ctx CommandContext, args []string) error {
	user, member := parseArgs(ctx, args)
	primTs, _ := discordgo.SnowflakeTimestamp(user.ID)
	servTs, _ := member.JoinedAt.Parse()
	pres, _ := ctx.Session.State.Presence(ctx.Guild.ID, user.ID)
	primTsS := strconv.Itoa(int(math.Ceil(time.Since(primTs).Hours() / 24)))
	servTsS := strconv.Itoa(int(math.Ceil(time.Since(servTs).Hours() / 24)))
	var game string
	var title string
	var roles string
	var desc string

	if member.Nick != "" {
		title = user.Username + "#" + user.Discriminator + " - " + member.Nick
	} else {
		title = user.Username + "#" + user.Discriminator
	}

	if pres.Game == nil {
		game = "nothing."
	} else {
		game = pres.Game.Name
	}

	desc = "Playing " + game
	if user.Bot {
		desc += "\n> Bot account"
	}

	if len(ctx.Member.Roles) > 0 {

		for i, r := range member.Roles {
			n, _ := ctx.Session.State.Role(ctx.Guild.ID, r)
			if i == 0 {
				roles += n.Name
			} else if i == len(member.Roles)-1 {
				roles += ", " + n.Name + "."
			} else {
				roles += ", " + n.Name
			}
		}
	}
	embedBuilder := Shared.NewEmbed()
	embedBuilder.SetTitle(title)
	embedBuilder.SetColor(Shared.COLOR)
	embedBuilder.SetThumbnail(user.AvatarURL("1024"))
	embedBuilder.SetDescription(desc)
	embedBuilder.AddInlineField("Joined Discord on", primTs.UTC().Format(tformat)+"\n**"+primTsS+"** days ago")
	embedBuilder.AddInlineField("Joined This Server on", servTs.UTC().Format(tformat)+"\n**"+servTsS+"** days ago")
	embedBuilder.AddField("Roles (**"+strconv.Itoa(len(member.Roles))+"**)", roles)
	embedBuilder.AddInlineField("Is Bot Owner?", strings.Title(strconv.FormatBool(ctx.Manager.IsOwner(user.ID))))
	_, err := ctx.ReplyEmbed(embedBuilder.MessageEmbed)
	return err
}

func parseArgs(ctx CommandContext, args []string) (*discordgo.User, *discordgo.Member) {
	if len(ctx.Message.Mentions) > 0 {
		u := ctx.Message.Mentions[0]
		m, _ := ctx.Session.State.Member(ctx.Guild.ID, u.ID)
		return u, m
	} else if len(args) > 0 {
		u, err := ctx.Session.User(args[0])
		if err != nil {
			return ctx.User, ctx.Member
		}

		m, err := ctx.Session.State.Member(ctx.Guild.ID, u.ID)
		if err != nil {
			return ctx.User, ctx.Member
		}

		return u, m

	}
	return ctx.User, ctx.Member
}
