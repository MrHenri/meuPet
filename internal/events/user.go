package events

import (
	"time"

	"github.com/MrHenri/meuPet/pkg/events"
)

const (
	UserCreated events.EventName = "user-created"
)

type UserEvent struct {
	Name    events.EventName
	Payload interface{}
}

func NewUserEvent(name events.EventName) *UserEvent {
	return &UserEvent{Name: name}
}

func (u *UserEvent) GetName() events.EventName {
	return u.Name
}

func (u *UserEvent) GetDateTime() time.Time {
	return time.Now()
}

func (u *UserEvent) GetPayload() interface{} {
	return u.Payload
}

func (u *UserEvent) SetPayload(payload interface{}) {
	u.Payload = payload
}
