package imbus

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/event"
)

var _ event.Bus = &bus{}

type bus struct {
	handlers map[event.Type][]event.Handler
}

func New() *bus {
	return &bus{}
}

func (b *bus) Publish(ctx context.Context, events []event.Event) error {
	for _, e := range events {
		handlers, ok := b.handlers[e.Type()]
		if ok {
			for _, handler := range handlers {
				if err := handler.Handle(ctx, e); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (b *bus) Subscribe(eventType event.Type, handler event.Handler) {
	handlers, ok := b.handlers[eventType]
	if !ok {
		handlers = []event.Handler{}
	}
	handlers = append(handlers, handler)
	b.handlers[eventType] = handlers
}
