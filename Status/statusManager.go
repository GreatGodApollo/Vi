package Status

import (
	"github.com/bwmarrin/discordgo"
	"github.com/greatgodapollo/Vi/Configuration"
	"math/rand"
	"time"
)

type StatusManager struct {
	Entries  []string
	Interval string
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

func NewStatusManager(c Configuration.Configuration) *StatusManager {
	return &StatusManager{
		Entries:  c.Bot.Statuses,
		Interval: c.Bot.StatusInterval,
	}
}

func (sm *StatusManager) OnReady(s *discordgo.Session, e *discordgo.Ready) {
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
