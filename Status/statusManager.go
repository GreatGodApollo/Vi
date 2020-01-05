package Status

import (
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Configuration"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type StatusManager struct {
	Entries  []string
	Interval string
	log      *logrus.Logger
}

func (s *StatusManager) AddEntry(entry string) {
	s.Entries = append(s.Entries, entry)
}

func (s *StatusManager) SetEntries(entries []string) {
	s.Entries = entries
}

func (s *StatusManager) SetInterval(interval string) {
	s.Interval = interval
}

func (sm *StatusManager) SetStatus(s *discordgo.Session) {
	i := rand.Intn(len(sm.Entries))
	_ = s.UpdateStatus(0, sm.Entries[i])
}

func NewStatusManager(c Configuration.Configuration, log *logrus.Logger) *StatusManager {
	return &StatusManager{
		Entries:  c.Bot.Statuses,
		Interval: c.Bot.StatusInterval,
		log:      log,
	}
}

func (sm *StatusManager) OnReady(s *discordgo.Session, e *discordgo.Ready) {
	sm.log.Info("Bot is ready to receive commands!")
	interval, err := time.ParseDuration(sm.Interval)
	if err != nil {
		return
	}

	sm.SetStatus(s)

	tick := time.NewTicker(interval)
	for range tick.C {
		sm.SetStatus(s)
	}
}
