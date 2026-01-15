# Handler Testing Standard

This document defines the **standard pattern** for testing HTTP handlers in the `photography-api` project. All developers and AI contributors must follow these guidelines to ensure consistency, reliability, and maintainability of handler tests.

---

## 1. Principles

1. **Isolate domain logic:** Use real use cases and handlers, but mock repositories and external dependencies.
2. **Enforce mock verification:** Every test must verify that all expected mock calls occurred.
3. **Minimize boilerplate:** Use helpers for router and mock setup.
4. **Automatic mock assertion:** Use `t.Cleanup` to always call `AssertExpectations` on mocks.
5. **Comprehensive scenarios:** Each handler test must cover:
   - Success with data
   - Success with empty results (if applicable)
   - Error from use case (e.g., validation, repository, or unexpected errors)
6. **Readable and DRY:** Test logic should be clear, concise, and avoid repetition.

---

## 2. File Structure

- Place handler tests next to their implementation:
  ```
  /internal/web/infrastructure/http/<handler>/handler_test.go
  ```
- Place reusable mocks in:
  ```
  /internal/web/test/mocks/
  ```
- Place test data builders (mothers) in:
  ```
  /internal/web/test/<entity>mother.go
  ```

---

## 3. Mock Repository with Auto Assertion

Always use a helper to create mocks and register `AssertExpectations` with `t.Cleanup`:

```go
func prepareMockWithAutoAssert(t *testing.T) *MockPhotoRepository {
    mockRepo := new(MockPhotoRepository)
    t.Cleanup(func() {
        mockRepo.AssertExpectations(t)
    })
    return mockRepo
}
```

- This ensures no test can skip mock verification.

---

## 4. Router and Handler Setup Helper

Centralize handler and router setup in a helper:

```go
func preparePhotoHandlerWithMock(mockRepo *MockPhotoRepository) *gin.Engine {
    router := gin.New()
    uc := application.NewSearchPhotosUseCase(mockRepo)
    h := handler.NewPhotoHandler(uc)
    router.GET("/photos", h.SearchPhotos)
    return router
}
```

- This keeps tests DRY and focused on logic, not setup.

---

## 5. Test Structure

Each handler test file must:

- Use `t.Parallel()` at the top level.
- Use `t.Run()` for each scenario.
- Use helpers for mock and router setup.
- Assert HTTP status, response body, and (optionally) headers.
- Rely on `t.Cleanup` for mock assertion.

**Example:**

```go
func TestPhotoHandler_SearchPhotos(t *testing.T) {
    t.Parallel()

    t.Run("Success with data", func(t *testing.T) {
        mockRepo := prepareMockWithAutoAssert(t)
        mockRepo.On("Search", mock.Anything).Return(photos, nil)
        router := preparePhotoHandlerWithMock(mockRepo)
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("GET", "/photos", nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, 200, w.Code)
        // ...assert response body...
    })

    t.Run("Error from use case", func(t *testing.T) {
        mockRepo := prepareMockWithAutoAssert(t)
        mockRepo.On("Search", mock.Anything).Return(nil, errors.New("fail"))
        router := preparePhotoHandlerWithMock(mockRepo)
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("GET", "/photos", nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, 500, w.Code)
        // ...assert error response...
    })
}
```

---

## 6. Adding New Scenarios

For each new scenario:

1. Add a new `t.Run("scenario name", ...)`.
2. Use the mock and router helpers.
3. Configure the mock for the scenario.
4. Make the HTTP request with `httptest`.
5. Assert status, body, and (if needed) headers.
6. Let `t.Cleanup` handle mock assertion.

---

## 7. Best Practices

- Always use subtests (`t.Run`) for each scenario.
- Use `Once()` on mocks to ensure exact call count.
- Keep mock setup, request execution, and assertions clearly separated.
- Never put domain logic in the handler or mock; tests should only orchestrate and verify.
- Use test data builders (mothers) for complex objects.
- Prefer table-driven tests for similar scenarios.

---

## 8. Benefits

- **Consistency:** All handler tests look and behave the same.
- **Safety:** No skipped mock assertions.
- **Readability:** Easy for humans and AI to understand and extend.
- **Scalability:** Easy to add new handlers and scenarios.
- **AI-friendly:** Clear, repeatable pattern for automated test generation.

---

## 9. Checklist for Handler Tests

- [ ] Uses mock repository with `t.Cleanup` for assertions
- [ ] Uses router/helper setup
- [ ] Covers success (with data), empty, and error scenarios
- [ ] Uses subtests for each scenario
- [ ] Asserts HTTP status and response body
- [ ] No domain logic in test or handler
- [ ] No skipped mock assertions

---

**All new and existing handler tests must follow this standard.**
```

This guideline will help ensure all handler tests in your application are robust, maintainable, and easy for both humans and AI to extend.

## User

return the format ass markdown to copy or write on original doc file

## Assistant

Absolutely! Here’s the improved guideline in Markdown format, ready for you to copy and use in your original `docs/handler.md` file:

```md
# Handler Testing Standard

This document defines the **standard pattern** for testing HTTP handlers in the `photography-api` project. All developers and AI contributors must follow these guidelines to ensure consistency, reliability, and maintainability of handler tests.

---

## 1. Principles

1. **Isolate domain logic:** Use real use cases and handlers, but mock repositories and external dependencies.
2. **Enforce mock verification:** Every test must verify that all expected mock calls occurred.
3. **Minimize boilerplate:** Use helpers for router and mock setup.
4. **Automatic mock assertion:** Use `t.Cleanup` to always call `AssertExpectations` on mocks.
5. **Comprehensive scenarios:** Each handler test must cover:
   - Success with data
   - Success with empty results (if applicable)
   - Error from use case (e.g., validation, repository, or unexpected errors)
6. **Readable and DRY:** Test logic should be clear, concise, and avoid repetition.

---

## 2. File Structure

- Place handler tests next to their implementation:
  ```
  /internal/web/infrastructure/http/<handler>/handler_test.go
  ```
- Place reusable mocks in:
  ```
  /internal/web/test/mocks/
  ```
- Place test data builders (mothers) in:
  ```
  /internal/web/test/<entity>mother.go
  ```

---

## 3. Mock Repository with Auto Assertion

Always use a helper to create mocks and register `AssertExpectations` with `t.Cleanup`:

```go
func prepareMockWithAutoAssert(t *testing.T) *MockPhotoRepository {
    mockRepo := new(MockPhotoRepository)
    t.Cleanup(func() {
        mockRepo.AssertExpectations(t)
    })
    return mockRepo
}
```

- This ensures no test can skip mock verification.

---

## 4. Router and Handler Setup Helper

Centralize handler and router setup in a helper:

```go
func preparePhotoHandlerWithMock(mockRepo *MockPhotoRepository) *gin.Engine {
    router := gin.New()
    uc := application.NewSearchPhotosUseCase(mockRepo)
    h := handler.NewPhotoHandler(uc)
    router.GET("/photos", h.SearchPhotos)
    return router
}
```

- This keeps tests DRY and focused on logic, not setup.

---

## 5. Test Structure

Each handler test file must:

- Use `t.Parallel()` at the top level.
- Use `t.Run()` for each scenario.
- Use helpers for mock and router setup.
- Assert HTTP status, response body, and (optionally) headers.
- Rely on `t.Cleanup` for mock assertion.

**Example:**

```go
func TestPhotoHandler_SearchPhotos(t *testing.T) {
    t.Parallel()

    t.Run("Success with data", func(t *testing.T) {
        mockRepo := prepareMockWithAutoAssert(t)
        mockRepo.On("Search", mock.Anything).Return(photos, nil)
        router := preparePhotoHandlerWithMock(mockRepo)
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("GET", "/photos", nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, 200, w.Code)
        // ...assert response body...
    })

    t.Run("Error from use case", func(t *testing.T) {
        mockRepo := prepareMockWithAutoAssert(t)
        mockRepo.On("Search", mock.Anything).Return(nil, errors.New("fail"))
        router := preparePhotoHandlerWithMock(mockRepo)
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("GET", "/photos", nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, 500, w.Code)
        // ...assert error response...
    })
}
```

---

## 6. Adding New Scenarios

For each new scenario:

1. Add a new `t.Run("scenario name", ...)`.
2. Use the mock and router helpers.
3. Configure the mock for the scenario.
4. Make the HTTP request with `httptest`.
5. Assert status, body, and (if needed) headers.
6. Let `t.Cleanup` handle mock assertion.

---

## 7. Best Practices

- Always use subtests (`t.Run`) for each scenario.
- Use `Once()` on mocks to ensure exact call count.
- Keep mock setup, request execution, and assertions clearly separated.
- Never put domain logic in the handler or mock; tests should only orchestrate and verify.
- Use test data builders (mothers) for complex objects.
- Prefer table-driven tests for similar scenarios.

---

## 8. Benefits

- **Consistency:** All handler tests look and behave the same.
- **Safety:** No skipped mock assertions.
- **Readability:** Easy for humans and AI to understand and extend.
- **Scalability:** Easy to add new handlers and scenarios.
- **AI-friendly:** Clear, repeatable pattern for automated test generation.

---

## 9. Checklist for Handler Tests

- [ ] Uses mock repository with `t.Cleanup` for assertions
- [ ] Uses router/helper setup
- [ ] Covers success (with data), empty, and error scenarios
- [ ] Uses subtests for each scenario
- [ ] Asserts HTTP status and response body
- [ ] No domain logic in test or handler
- [ ] No skipped mock assertions

---

**All new and existing handler tests must follow this standard.**
```
