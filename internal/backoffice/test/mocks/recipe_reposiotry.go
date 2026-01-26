package mocks

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
	"github.com/stretchr/testify/mock"
)

type MockRecipeRepository struct {
	mock.Mock
}

type repo interface {
	recipe.Repository
	recipe.Repository
}

var _ repo = &MockRecipeRepository{}

func (m *MockRecipeRepository) Exists(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(bool), args.Error(1)
}

func (m *MockRecipeRepository) Search(ctx context.Context, criteria recipe.Criteria) ([]*recipe.Recipe, error) {
	args := m.Called(ctx, criteria)
	return args.Get(0).([]*recipe.Recipe), args.Error(1)
}

func (m *MockRecipeRepository) Save(ctx context.Context, recipe *recipe.Recipe) error {
	args := m.Called(ctx, recipe)
	return args.Error(0)
}
