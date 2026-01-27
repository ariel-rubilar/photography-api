package searchphotos_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/middleware"
	"github.com/ariel-rubilar/photography-api/internal/web/test/mocks"
	"github.com/ariel-rubilar/photography-api/internal/web/test/photodtomother"
	sharedmocks "github.com/ariel-rubilar/photography-api/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ariel-rubilar/photography-api/internal/shared/domain/domainerror"
	"github.com/ariel-rubilar/photography-api/internal/web/infrastructure/http/searchphotos"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/photoquery"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
)

type Providers struct {
	Repo *mocks.MockPhotoRepository
}

func prepareMockWithAutoAssert(t *testing.T) Providers {
	mockRepo := new(mocks.MockPhotoRepository)

	t.Cleanup(func() {
		mockRepo.AssertExpectations(t)
	})

	return Providers{
		Repo: mockRepo,
	}
}

func preparePhotoHandlerWithProviders(providers Providers) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	uc := searcher.New(providers.Repo)
	h := searchphotos.NewHandler(uc)

	logger := sharedmocks.NewNoOpLogger()

	router.Use(
		middleware.ErrorHandler(logger),
	)

	router.GET("/photos", h)
	return router
}

func TestPhotoHandler_SearchPhotos(t *testing.T) {
	t.Parallel()

	t.Run("Success with empty data", func(t *testing.T) {

		providers := prepareMockWithAutoAssert(t)

		router := preparePhotoHandlerWithProviders(providers)

		photos := photodtomother.NewPhotoDTOList(0)

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/photos", nil)
		require.NoError(t, err)

		providers.Repo.On("Search", req.Context(), photoquery.Criteria{}).Return(photos, nil).Once()

		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var resp struct {
			Data []searchphotos.PhotoDTO `json:"data"`
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data, 0)

	})

	t.Run("Success with data", func(t *testing.T) {

		providers := prepareMockWithAutoAssert(t)

		router := preparePhotoHandlerWithProviders(providers)
		photos := photodtomother.NewPhotoDTOList(2)

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/photos", nil)
		require.NoError(t, err)

		providers.Repo.On("Search", req.Context(), photoquery.Criteria{}).Return(photos, nil).Once()

		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		var resp struct {
			Data []searchphotos.PhotoDTO `json:"data"`
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data, 2)

		for i, photo := range photos {
			actual := resp.Data[i]

			primitives := photo

			expected := searchphotos.PhotoDTO{
				ID:    primitives.ID,
				Title: primitives.Title,
				URL:   primitives.URL,
				Recipe: searchphotos.PhotoRecipe{
					Name: primitives.Recipe.Name,
					Settings: searchphotos.PhotoRecipeSettings{
						FilmSimulation:       primitives.Recipe.Settings.FilmSimulation,
						DynamicRange:         primitives.Recipe.Settings.DynamicRange,
						Highlight:            primitives.Recipe.Settings.Highlight,
						Shadow:               primitives.Recipe.Settings.Shadow,
						Color:                primitives.Recipe.Settings.Color,
						NoiseReduction:       primitives.Recipe.Settings.NoiseReduction,
						Sharpening:           primitives.Recipe.Settings.Sharpening,
						Clarity:              primitives.Recipe.Settings.Clarity,
						GrainEffect:          primitives.Recipe.Settings.GrainEffect,
						ColorChromeEffect:    primitives.Recipe.Settings.ColorChromeEffect,
						ColorChromeBlue:      primitives.Recipe.Settings.ColorChromeBlue,
						WhiteBalance:         primitives.Recipe.Settings.WhiteBalance,
						Iso:                  primitives.Recipe.Settings.Iso,
						ExposureCompensation: primitives.Recipe.Settings.ExposureCompensation,
					},
					Link: primitives.Recipe.Link,
				},
			}
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("Error searching photos", func(t *testing.T) {

		providers := prepareMockWithAutoAssert(t)

		router := preparePhotoHandlerWithProviders(providers)

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/photos", nil)
		require.NoError(t, err)

		providers.Repo.On("Search", req.Context(), photoquery.Criteria{}).Return([]*photoquery.PhotoDTO{}, domainerror.Validation{
			Reason: "TEST",
		}).Once()

		router.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)

		var resp sharedhttp.ErrorResponse
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t, "VALIDATION_ERROR", resp.Error.Code)
		assert.Equal(t, "TEST", resp.Error.Message)

	})
}
