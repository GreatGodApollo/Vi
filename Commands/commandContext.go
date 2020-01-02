package Commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Status"
	"io"
)

type CommandContext struct {
	Session       *discordgo.Session
	Event         *discordgo.MessageCreate
	Manager       *CommandManager
	StatusManager *Status.StatusManager
	Message       *discordgo.Message
	User          *discordgo.User
	Channel       *discordgo.Channel
	Guild         *discordgo.Guild
	Member        *discordgo.Member
}

func (ctx *CommandContext) Reply(message string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSend(ctx.Channel.ID, message)
}

func (ctx *CommandContext) ReplyEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendEmbed(ctx.Channel.ID, embed)
}

func (ctx *CommandContext) ReplyFile(filename string, file io.Reader) (*discordgo.Message, error) {
	return ctx.Session.ChannelFileSend(ctx.Channel.ID, filename, file)
}
