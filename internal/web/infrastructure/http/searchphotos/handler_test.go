package searchphotos_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/middleware"
	"github.com/ariel-rubilar/photography-api/internal/web/photo"
	"github.com/ariel-rubilar/photography-api/internal/web/test/photomother"
	"github.com/ariel-rubilar/photography-api/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	domainerror "github.com/ariel-rubilar/photography-api/internal/shared/domain/error"
	"github.com/ariel-rubilar/photography-api/internal/web/infrastructure/http/searchphotos"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
)

type mockRepository struct {
	Data *[]*photo.Photo
	Err  error
}

func (m *mockRepository) Search(ctx context.Context) ([]*photo.Photo, error) {
	if m.Err != nil {
		return []*photo.Photo{}, m.Err
	}
	return *m.Data, nil
}

func TestHandler_Success_Response(t *testing.T) {

	gin.SetMode(gin.TestMode)
	photos := photomother.NewPhotoList(2)

	var r photo.Repository = &mockRepository{
		Data: &photos,
	}

	uc := searcher.New(r)
	h := searchphotos.NewHandler(uc)

	router := gin.New()

	logger := mocks.NewNoOpLogger()

	router.Use(
		middleware.ErrorHandler(logger),
	)

	router.GET("/test", h)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var resp struct {
		Data []searchphotos.PhotoDTO `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data, 2)

	for i, photo := range photos {
		actual := resp.Data[i]

		primitives := photo.ToPrimitives()

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
}

func TestHandler_Error_Response(t *testing.T) {

	gin.SetMode(gin.TestMode)

	var r photo.Repository = &mockRepository{
		Err: domainerror.Validation{
			Reason: "TEST",
		},
	}

	uc := searcher.New(r)
	h := searchphotos.NewHandler(uc)

	router := gin.New()

	logger := mocks.NewNoOpLogger()

	router.Use(
		middleware.ErrorHandler(logger),
	)

	router.GET("/test", h)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)

	var resp sharedhttp.ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "VALIDATION_ERROR", resp.Error.Code)
	assert.Equal(t, "TEST", resp.Error.Message)

}
