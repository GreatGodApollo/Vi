package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Commands"
	"github.com/greatgodapollo/Vi/Configuration"
	"github.com/greatgodapollo/Vi/Status"
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

	// Create discordgo client
	client, err := discordgo.New("Bot " + Config.Bot.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the CommandManager
	sm := Status.NewStatusManager(Config, log)
	cmdm := Commands.NewCommandManager(Config, sm, log, true)

	// Add the commands
	cmdm.AddCommand(Commands.NewAboutCommand())
	cmdm.AddCommand(Commands.NewHelpCommand())
	cmdm.AddCommand(Commands.NewOwnerCommand())
	cmdm.AddCommand(Commands.NewPingCommand())
	cmdm.AddCommand(Commands.NewTagCommand())

	// Load the tags file (will error if does not exist)
	Commands.LoadTags("tags.json", log)

	// Add the command handler
	client.AddHandler(cmdm.CommandHandler)

	// Add the StatusHandler
	client.AddHandler(sm.OnReady)

	// Connect to websocket and begin listening
	err = client.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmdm.AddPrefix("<@!" + client.State.User.ID + "> ")
	cmdm.AddPrefix("<@!" + client.State.User.ID + ">")

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
