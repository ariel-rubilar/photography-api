package mocks

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/web/photo"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/photoquery"
	"github.com/stretchr/testify/mock"
)

type MockPhotoRepository struct {
	mock.Mock
}

type repo interface {
	photo.Repository
	photoquery.Repository
}

var _ repo = &MockPhotoRepository{}

func (m *MockPhotoRepository) Search(ctx context.Context, criterial photoquery.Criteria) ([]*photoquery.PhotoDTO, error) {
	args := m.Called(ctx, criterial)
	return args.Get(0).([]*photoquery.PhotoDTO), args.Error(1)
}

func (m *MockPhotoRepository) Save(ctx context.Context, new *photo.Photo) error {
	args := m.Called(ctx, new)
	return args.Error(0)
}

func (m *MockPhotoRepository) Exists(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}
