package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Bus interface {
	Publish(context.Context, []Event) error
	Subscribe(Type, Handler)
}

type Handler interface {
	Handle(context.Context, Event) error
}

type Type string

type Event interface {
	ID() string
	OccurredOn() time.Time
	Type() Type
}

type BaseEvent struct {
	id         string
	occurredOn time.Time
}

func NewBaseEvent() BaseEvent {
	return BaseEvent{
		id:         uuid.New().String(),
		occurredOn: time.Now(),
	}
}

func (b BaseEvent) ID() string {
	return b.id
}

func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}
