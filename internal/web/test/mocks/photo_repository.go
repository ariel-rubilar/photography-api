package mocks

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/web/photo"
	"github.com/stretchr/testify/mock"
)

type MockPhotoRepository struct {
	mock.Mock
}

var _ photo.Repository = &MockPhotoRepository{}

func (m *MockPhotoRepository) Search(ctx context.Context) ([]*photo.Photo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*photo.Photo), args.Error(1)
}

func (m *MockPhotoRepository) Save(ctx context.Context, new *photo.Photo) error {
	args := m.Called(ctx, new)
	return args.Error(0)
}

func (m *MockPhotoRepository) Exists(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}
