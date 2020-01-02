package Events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (eventm *EventManager) OnGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	if e.Guild.Unavailable {
		return
	}

	channel := e.Channels[0]

	embed := &discordgo.MessageEmbed{
		Title:       "Hi, My name is Vi!",
		Description: fmt.Sprintf("Hello there! I'm your new friend Vi. To check out what I can do, go ahead and run `%shelp`!", eventm.Config.Bot.Prefixes[0]),
		Color:       0,
	}

	_, _ = s.ChannelMessageSendEmbed(channel.ID, embed)
}
