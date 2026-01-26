package mocks

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photoquery"
	"github.com/stretchr/testify/mock"
)

type MockPhotoQueryRepository struct {
	mock.Mock
}

var _ photoquery.Repository = &MockPhotoQueryRepository{}

func (m *MockPhotoQueryRepository) Search(ctx context.Context, criteria photoquery.Criteria) ([]*photoquery.PhotoDTO, error) {
	args := m.Called(ctx, criteria)
	return args.Get(0).([]*photoquery.PhotoDTO), args.Error(1)
}
