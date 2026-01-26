package searchrecipes_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipequery"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/http/searchrecipes"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/test/mocks"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/test/mocks/recipedtomother"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/middleware"
	sharedmocks "github.com/ariel-rubilar/photography-api/test/mocks"
	"github.com/gin-gonic/gin"
)

type Providers struct {
	Repo *mocks.MockRecipeQueryRepository
}

func prepareMockWithAutoAssert(t *testing.T) Providers {
	mockRepo := new(mocks.MockRecipeQueryRepository)

	t.Cleanup(func() {
		mockRepo.AssertExpectations(t)
	})

	return Providers{
		Repo: mockRepo,
	}
}

func prepareHandlerWithProviders(providers Providers) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	uc := recipesearcher.New(providers.Repo)
	h := searchrecipes.NewHandler(uc)

	logger := sharedmocks.NewNoOpLogger()

	router.Use(
		middleware.ErrorHandler(logger),
	)
	router.GET("/recipes", h)

	return router
}

func TestSearchRecipesHandler(t *testing.T) {
	t.Parallel()

	t.Run("error searching recipes", func(t *testing.T) {

		providers := prepareMockWithAutoAssert(t)
		router := prepareHandlerWithProviders(providers)

		req := httptest.NewRequest("GET", "/recipes", nil)
		w := httptest.NewRecorder()

		providers.Repo.On("Search", req.Context(), recipequery.Criteria{}).
			Return(recipedtomother.NewRecipeDTOList(0), errors.New("internal server error")).
			Once()

		router.ServeHTTP(w, req)

		var response sharedhttp.ErrorResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "INTERNAL_ERROR", response.Error.Code)
		assert.Equal(t, "internal server error", response.Error.Message)
	})

	t.Run("successfully search recipes", func(t *testing.T) {

		providers := prepareMockWithAutoAssert(t)
		router := prepareHandlerWithProviders(providers)

		req := httptest.NewRequest("GET", "/recipes", nil)
		w := httptest.NewRecorder()

		recipes := recipedtomother.NewRecipeDTOList(2)

		providers.Repo.On("Search", req.Context(), recipequery.Criteria{}).
			Return(recipes, nil).
			Once()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response struct {
			Data []searchrecipes.RecipeDTO `json:"data"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Len(t, response.Data, 2)

		for _, recipe := range response.Data {
			expected := searchrecipes.RecipeDTO{
				ID:   recipe.ID,
				Name: recipe.Name,
				Settings: searchrecipes.SettingDTO{
					FilmSimulation:       recipe.Settings.FilmSimulation,
					DynamicRange:         recipe.Settings.DynamicRange,
					Highlight:            recipe.Settings.Highlight,
					Shadow:               recipe.Settings.Shadow,
					Color:                recipe.Settings.Color,
					NoiseReduction:       recipe.Settings.NoiseReduction,
					Sharpening:           recipe.Settings.Sharpening,
					Clarity:              recipe.Settings.Clarity,
					GrainEffect:          recipe.Settings.GrainEffect,
					ColorChromeEffect:    recipe.Settings.ColorChromeEffect,
					ColorChromeBlue:      recipe.Settings.ColorChromeBlue,
					WhiteBalance:         recipe.Settings.WhiteBalance,
					Iso:                  recipe.Settings.Iso,
					ExposureCompensation: recipe.Settings.ExposureCompensation,
				},
				Link: recipe.Link,
			}

			assert.Equal(t, expected, recipe)
		}
	})
}
