# PhotoHandler Testing Pattern

This document explains the **recommended pattern** for testing HTTP endpoints in the `photography-api` project using **Gin, Testify, and DDD/Hexagonal architecture**. It ensures consistency and prevents common mistakes, like forgetting to assert mock expectations.

---

## 1. Principles

1. **Use real use cases and handlers**, but **mock repositories** to isolate domain logic.
2. Each test must **verify repository calls** to ensure correct behavior.
3. Avoid repeating boilerplate: use **helpers** for router and mock setup.
4. **Automatic mock assertion**: every mock must call `AssertExpectations` even if the developer forgets it.
5. Test scenarios should include:

   * Success with data
   * Success with empty results (optional)
   * Error from use case

---

## 2. File Structure

```
/internal/web/adapter/http/handler/photo_handler_test.go
/internal/web/test/mock_repository.go (optional, reusable mocks)
```

* Tests live next to handlers.
* Mocks can be centralized if reused across tests.

---

## 3. Mock Repository with Auto Assertion

Use `t.Cleanup` to ensure `AssertExpectations` is called automatically:

```go
func prepareMockWithAutoAssert(t *testing.T) *MockPhotoRepository {
    mockRepo := new(MockPhotoRepository)
    t.Cleanup(func() {
        mockRepo.AssertExpectations(t)
    })
    return mockRepo
}
```

* Ensures expectations are validated at the **end of each test**.
* Developers cannot skip this step.

---

## 4. Router Helper

Centralize handler and router setup:

```go
func preparePhotoHandlerWithMock(mockRepo *MockPhotoRepository) *gin.Engine {
    router := gin.Default()
    uc := application.NewSearchPhotosUseCase(mockRepo)
    h := handler.NewPhotoHandler(uc)
    router.GET("/photos", h.SearchPhotos)
    return router
}
```

* Returns a ready-to-use `gin.Engine`.
* Keeps test code DRY.

---

## 5. Example Test File

```go
func TestPhotoHandler_SearchPhotos(t *testing.T) {
    t.Parallel()

    t.Run("success with data", func(t *testing.T) {
        mockRepo := prepareMockWithAutoAssert(t)
        mockRepo.On("Search", mock.Anything, "sunset").
            Return([]domain.Photo{{ID: "1", Title: "Sunset"}}, nil).
            Once()

        router := preparePhotoHandlerWithMock(mockRepo)

        w := httptest.NewRecorder()
        req, _ := http.NewRequest(http.MethodGet, "/photos?query=sunset", nil)
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusOK, w.Code)

        var resp []domain.Photo
        _ = json.Unmarshal(w.Body.Bytes(), &resp)
        assert.Len(t, resp, 1)
        assert.Equal(t, "Sunset", resp[0].Title)
    })

    t.Run("error executing use case", func(t *testing.T) {
        mockRepo := prepareMockWithAutoAssert(t)
        mockRepo.On("Search", mock.Anything, "error-query").
            Return(nil, errors.New("db error"))
            Once()

        router := preparePhotoHandlerWithMock(mockRepo)

        w := httptest.NewRecorder()
        req, _ := http.NewRequest(http.MethodGet, "/photos?query=error-query", nil)
        router.ServeHTTP(w, req)

        assert.Equal(t, http.StatusInternalServerError, w.Code)
    })
}
```

---

## 6. How to Add New Scenarios

1. Use a new `t.Run()` subtest.
2. Prepare a mock using `prepareMockWithAutoAssert`.
3. Configure the repository behavior for the scenario.
4. Use the router helper to create the endpoint.
5. Call the endpoint via `httptest` and assert:

   * HTTP status
   * Body (optional)
   * Repository calls are automatically verified

---

## 7. Benefits of This Pattern

* **Consistency:** all developers follow the same structure.
* **Safety:** no scenario can skip mock verification.
* **Readability:** test logic (request + assertions) is visible.
* **Scalability:** easy to add new scenarios without repeating boilerplate.
* **AI-friendly:** pattern is clear for automatic test generation or suggestions.

---

## 8. Recommended Best Practices

* Always use **subtests (`t.Run`)** for each scenario.
* Keep mock setup separate from endpoint call and assertions.
* Use **`Once()`** on mocks to ensure exact call count.
* Use **`t.Cleanup`** for automatic assertion of expectations.
* Avoid putting domain logic inside handler or mock; the test only orchestrates behavior.
