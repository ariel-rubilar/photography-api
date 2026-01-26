package mocks

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipequery"
	"github.com/stretchr/testify/mock"
)

type MockRecipeQueryRepository struct {
	mock.Mock
}

var _ recipequery.Repository = &MockRecipeQueryRepository{}

func (m *MockRecipeQueryRepository) Search(ctx context.Context, criteria recipequery.Criteria) ([]*recipequery.RecipeDTO, error) {
	args := m.Called(ctx, criteria)
	return args.Get(0).([]*recipequery.RecipeDTO), args.Error(1)
}
