package mocks

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/stretchr/testify/mock"
)

var _ photo.Repository = &MockPhotoRepository{}

type MockPhotoRepository struct {
	mock.Mock
}

func (m *MockPhotoRepository) Save(ctx context.Context, p *photo.Photo) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}
