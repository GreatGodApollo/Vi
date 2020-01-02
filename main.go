package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Commands"
	"github.com/greatgodapollo/Vi/Configuration"
	"github.com/greatgodapollo/Vi/Status"
	"os"
	"os/signal"
	"syscall"
)

var Config Configuration.Configuration

func main() {

	// Load in configuration file
	Config = Configuration.LoadConfiguration("config.json")

	// Create discordgo client
	client, err := discordgo.New("Bot " + Config.Bot.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the CommandManager
	sm := Status.NewStatusManager(Config)
	cmdm := Commands.NewCommandManager(Config, sm, true)

	// Add the commands
	cmdm.AddCommand(Commands.NewHelpCommand())
	cmdm.AddCommand(Commands.NewPingCommand())
	cmdm.AddCommand(Commands.NewAboutCommand())

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

	fmt.Println("Bot now running. CTRL-C to exit.")

	// Wait until a term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close after term signal
	_ = client.Close()
}
