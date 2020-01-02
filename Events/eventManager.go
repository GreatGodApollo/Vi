package Events

import "github.com/greatgodapollo/Vi/Configuration"

type EventManager struct {
	Config Configuration.Configuration
}

func NewEventManager(c Configuration.Configuration) EventManager {
	return EventManager{c}
}
