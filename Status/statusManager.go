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
