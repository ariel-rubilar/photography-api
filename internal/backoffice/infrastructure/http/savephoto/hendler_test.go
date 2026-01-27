package savephoto_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/http/savephoto"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/test/mocks"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/test/photomother"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/middleware"
	sharedmocks "github.com/ariel-rubilar/photography-api/test/mocks"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type Providers struct {
	Repo       *mocks.MockPhotoRepository
	EventBus   *sharedmocks.MockEventBus
	RecipeRepo *mocks.MockRecipeRepository
}

func prepareMockWithAutoAssert(t *testing.T) Providers {
	mockRepo := new(mocks.MockPhotoRepository)
	mockEventBus := new(sharedmocks.MockEventBus)
	mockRecipeRepo := new(mocks.MockRecipeRepository)

	t.Cleanup(func() {
		mockRepo.AssertExpectations(t)
		mockEventBus.AssertExpectations(t)
		mockRecipeRepo.AssertExpectations(t)
	})

	return Providers{
		Repo:       mockRepo,
		EventBus:   mockEventBus,
		RecipeRepo: mockRecipeRepo,
	}
}

func prepareSavePhotoHandlerWithProviders(providers Providers) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	uc := photosaver.New(providers.Repo, providers.RecipeRepo, providers.EventBus)
	h := savephoto.NewHandler(uc)

	logger := sharedmocks.NewNoOpLogger()

	router.Use(
		middleware.ErrorHandler(logger),
	)
	router.POST("/photos", h)

	return router
}

func TestSavePhotoHandler(t *testing.T) {
	t.Parallel()

	t.Run("successfully creates photo", func(t *testing.T) {
		providers := prepareMockWithAutoAssert(t)
		router := prepareSavePhotoHandlerWithProviders(providers)

		expectedPhoto := photomother.NewPhoto()
		primitives := expectedPhoto.ToPrimitives()

		dto := &savephoto.PhotoDTO{
			Title:    primitives.Title,
			ID:       primitives.ID,
			URL:      primitives.URL,
			RecipeID: primitives.RecipeID,
		}

		body, err := json.Marshal(dto)

		req, err := http.NewRequest("POST", "/photos", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		providers.Repo.On("Save",
			req.Context(),
			mock.MatchedBy(func(actual *photo.Photo) bool {
				require.Equal(t, expectedPhoto.ToPrimitives(), actual.ToPrimitives())
				return true
			}),
		).Return(nil).Once()

		providers.Repo.On("Exists",
			req.Context(),
			expectedPhoto.ToPrimitives().ID,
		).Return(false, nil).Once()

		providers.EventBus.On("Publish", req.Context(), mock.MatchedBy(func(events []event.Event) bool {

			require.Len(t, events, 1)

			actual, ok := events[0].(photo.PhotoCreatedEvent)
			if !ok {
				return false
			}

			require.Equal(t, primitives.ID, actual.PhotoID())
			require.Equal(t, primitives.RecipeID, actual.RecipeID())
			require.Equal(t, primitives.ID, actual.PhotoID())
			require.Equal(t, photo.PhotoCreatedEventType, actual.Type())

			return true
		})).Return(nil).Once()

		providers.RecipeRepo.On("Exists", req.Context(), primitives.RecipeID).Return(true, nil).Once()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response struct {
			Data string `json:"data"`
		}

		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, "photo saved successfully", response.Data)
	})

	t.Run("error creating photo", func(t *testing.T) {
		providers := prepareMockWithAutoAssert(t)
		router := prepareSavePhotoHandlerWithProviders(providers)

		expectedPhoto := photomother.NewPhoto()
		primitives := expectedPhoto.ToPrimitives()

		dto := &savephoto.PhotoDTO{
			Title:    primitives.Title,
			ID:       primitives.ID,
			URL:      primitives.URL,
			RecipeID: primitives.RecipeID,
		}

		body, err := json.Marshal(dto)

		req, err := http.NewRequest("POST", "/photos", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		providers.RecipeRepo.On("Exists", req.Context(), primitives.RecipeID).Return(true, nil).Once()

		providers.Repo.On("Exists",
			req.Context(),
			expectedPhoto.ToPrimitives().ID,
		).Return(false, nil).Once()

		providers.Repo.On("Save",
			req.Context(),
			mock.MatchedBy(func(actual *photo.Photo) bool {
				require.Equal(t, expectedPhoto.ToPrimitives(), actual.ToPrimitives())
				return true
			}),
		).Return(errors.New("photo already exists")).Once()

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response sharedhttp.ErrorResponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, "INTERNAL_ERROR", response.Error.Code)
		assert.Equal(t, "internal server error", response.Error.Message)
	})

	t.Run("photo already exists", func(t *testing.T) {
		providers := prepareMockWithAutoAssert(t)
		router := prepareSavePhotoHandlerWithProviders(providers)

		expectedPhoto := photomother.NewPhoto()
		primitives := expectedPhoto.ToPrimitives()

		dto := &savephoto.PhotoDTO{
			Title:    primitives.Title,
			ID:       primitives.ID,
			URL:      primitives.URL,
			RecipeID: primitives.RecipeID,
		}

		body, err := json.Marshal(dto)

		req, err := http.NewRequest("POST", "/photos", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		providers.Repo.On("Exists",
			req.Context(),
			expectedPhoto.ToPrimitives().ID,
		).Return(true, nil).Once()

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)

		var response sharedhttp.ErrorResponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, "CONFLICT", response.Error.Code)
		assert.Equal(t, fmt.Sprintf("%s already exists", expectedPhoto.ToPrimitives().ID), response.Error.Message)
	})

	t.Run("recipe not found", func(t *testing.T) {
		providers := prepareMockWithAutoAssert(t)
		router := prepareSavePhotoHandlerWithProviders(providers)

		expectedPhoto := photomother.NewPhoto()
		primitives := expectedPhoto.ToPrimitives()

		dto := &savephoto.PhotoDTO{
			Title:    primitives.Title,
			ID:       primitives.ID,
			URL:      primitives.URL,
			RecipeID: primitives.RecipeID,
		}

		body, err := json.Marshal(dto)

		req, err := http.NewRequest("POST", "/photos", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		providers.Repo.On("Exists",
			req.Context(),
			expectedPhoto.ToPrimitives().ID,
		).Return(false, nil).Once()

		providers.RecipeRepo.On("Exists", req.Context(), primitives.RecipeID).Return(false, nil).Once()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response sharedhttp.ErrorResponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, "NOT_FOUND", response.Error.Code)
		assert.Equal(t, "recipe not found", response.Error.Message)
	})

	t.Run("error publish event", func(t *testing.T) {
		providers := prepareMockWithAutoAssert(t)
		router := prepareSavePhotoHandlerWithProviders(providers)

		expectedPhoto := photomother.NewPhoto()
		primitives := expectedPhoto.ToPrimitives()

		dto := &savephoto.PhotoDTO{
			Title:    primitives.Title,
			ID:       primitives.ID,
			URL:      primitives.URL,
			RecipeID: primitives.RecipeID,
		}

		body, err := json.Marshal(dto)

		req, err := http.NewRequest("POST", "/photos", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)
		providers.RecipeRepo.On("Exists", req.Context(), primitives.RecipeID).Return(true, nil).Once()

		providers.Repo.On("Save",
			req.Context(),
			mock.MatchedBy(func(actual *photo.Photo) bool {
				require.Equal(t, expectedPhoto.ToPrimitives(), actual.ToPrimitives())
				return true
			}),
		).Return(nil).Once()

		providers.Repo.On("Exists",
			req.Context(),
			expectedPhoto.ToPrimitives().ID,
		).Return(false, nil).Once()

		providers.EventBus.On("Publish", req.Context(), mock.MatchedBy(func(events []event.Event) bool {

			require.Len(t, events, 1)

			actual, ok := events[0].(photo.PhotoCreatedEvent)
			if !ok {
				return false
			}

			require.Equal(t, primitives.ID, actual.PhotoID())
			require.Equal(t, primitives.RecipeID, actual.RecipeID())
			require.Equal(t, primitives.ID, actual.PhotoID())
			require.Equal(t, photo.PhotoCreatedEventType, actual.Type())

			return true
		})).Return(errors.New("error")).Once()

		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response sharedhttp.ErrorResponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, "INTERNAL_ERROR", response.Error.Code)
		assert.Equal(t, "internal server error", response.Error.Message)
	})
}
