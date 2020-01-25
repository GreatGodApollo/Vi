//go:generate goversioninfo
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

package main

import (
	"fmt"
	"github.com/GreatGodApollo/Vi/Commands"
	"github.com/GreatGodApollo/Vi/Configuration"
	"github.com/GreatGodApollo/Vi/Shared"
	"github.com/GreatGodApollo/Vi/Status"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Config Configuration.Configuration
var log = logrus.New()

func main() {

	// Load in configuration file
	Config = Configuration.LoadConfiguration("config.json", log)

	// Create Logger
	if Config.Miscellaneous.ColorEnabled {
		log.SetFormatter(&prefixed.TextFormatter{
			ForceColors:     true,
			ForceFormatting: true,
			FullTimestamp:   true,
			TimestampFormat: time.RFC822Z,
		})
	} else {
		log.SetFormatter(&prefixed.TextFormatter{
			ForceColors:     false,
			ForceFormatting: true,
			FullTimestamp:   true,
			TimestampFormat: time.RFC822Z,
		})
	}
	log.Level = logrus.DebugLevel
	log.Info(fmt.Sprintf("Starting Vi %s", Shared.VERSION))

	// Create discordgo client
	client, err := discordgo.New("Bot " + Config.Bot.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the CommandManager
	sm := Status.NewStatusManager(Config, log)
	cmdm := Commands.NewCommandManager(Config, sm, log, true, CommandErrorFunc)

	// Register the commands
	registerCommands(cmdm)

	// Load the tags file (will error if does not exist)
	Commands.LoadTags("tags.json", log)

	// Add the command handler
	client.AddHandler(cmdm.CommandHandler)

	// Add the StatusHandler
	client.AddHandler(sm.OnReady)

	// Connect to websocket and begin listening
	err = client.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	cmdm.AddPrefix("<@!" + client.State.User.ID + "> ")
	cmdm.AddPrefix("<@!" + client.State.User.ID + ">")

	log.Info(fmt.Sprintf("Invite me at: https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot&permissions=379976", client.State.User.ID))

	// Wait until a term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close after term signal
	_ = client.Close()
}

func CommandErrorFunc(cmdm *Commands.CommandManager, err error) {
	cmdm.Logger.Error(err)
}

func registerCommands(cmdm Commands.CommandManager) {
	cmdm.AddCommand(Commands.NewAboutCommand())
	cmdm.AddCommand(Commands.NewHelpCommand())
	cmdm.AddCommand(Commands.NewInviteCommand())
	cmdm.AddCommand(Commands.NewOwnerCommand())
	cmdm.AddCommand(Commands.NewPingCommand())
	cmdm.AddCommand(Commands.NewTagCommand())
}
