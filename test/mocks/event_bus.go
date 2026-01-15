package mocks

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	"github.com/stretchr/testify/mock"
)

var _ event.Bus = &MockEventBus{}

type MockEventBus struct {
	mock.Mock
}

func (m *MockEventBus) Publish(ctx context.Context, events []event.Event) error {
	args := m.Called(ctx, events)
	return args.Error(0)
}

func (m *MockEventBus) Subscribe(eventType event.Type, handler event.Handler) {
	m.Called(eventType, handler)
}
