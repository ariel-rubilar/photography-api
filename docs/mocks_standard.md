# Mock Standards for Test Doubles

This document defines the **standard pattern** for creating and documenting mock implementations of interfaces in the `photography-api` project.  
All contributors—human and AI—must follow these guidelines to ensure consistency, reliability, and maintainability of test code.

---

## 1. Purpose of Mocks

- **Isolation:** Mocks allow you to test business logic, use cases, and handlers without relying on real infrastructure or external systems.
- **Verification:** Mocks support expectation setting and assertion, ensuring your code interacts with dependencies as expected.
- **Flexibility:** Mocks can be configured to return any data or error needed for your test scenarios.

---

## 2. File Location & Naming

- Place all mocks in the appropriate test-only package, typically:
  ```
  /internal/<context>/test/mocks/<interface>_mock.go
  ```
- The mock struct should be named `Mock<InterfaceName>`, e.g., `MockPhotoRepository`.

---

## 3. Implementation Pattern

- All mocks must embed `github.com/stretchr/testify/mock.Mock`.
- All methods of the interface must be implemented, using the `m.Called(...)` pattern.
- Always assert interface compliance with:
  ```go
  var _ <package>.<Interface> = &Mock<InterfaceName>{}
  ```

**Example:**

```go
package mocks

import (
    "context"
    "github.com/ariel-rubilar/photography-api/internal/web/photo"
    "github.com/stretchr/testify/mock"
)

// MockPhotoRepository is a testify-based mock for the photo.Repository interface.
// See /docs/mocks_standard.md for usage and conventions.
type MockPhotoRepository struct {
    mock.Mock
}

var _ photo.Repository = &MockPhotoRepository{}

func (m *MockPhotoRepository) Search(ctx context.Context) ([]*photo.Photo, error) {
    args := m.Called(ctx)
    return args.Get(0).([]*photo.Photo), args.Error(1)
}
```

---

## 4. Usage in Tests

- Create the mock in your test:
  ```go
  mockRepo := new(mocks.MockPhotoRepository)
  ```
- Set up expectations and return values:
  ```go
  mockRepo.On("Search", mock.Anything).Return(expectedPhotos, nil).Once()
  ```
- Always assert expectations:
  ```go
  t.Cleanup(func() { mockRepo.AssertExpectations(t) })
  // or
  defer mockRepo.AssertExpectations(t)
  ```

---

## 5. Documentation for Each Mock

Each mock file should include a short doc comment at the top, e.g.:

```go
// MockPhotoRepository is a testify-based mock for the photo.Repository interface.
// See /docs/mocks_standard.md for usage and conventions.
```

---

## 6. AI Guidelines for Generating Mocks

- Always follow the file structure and naming conventions above.
- Implement all interface methods using `m.Called(...)`.
- Assert interface compliance.
- Add a doc comment referencing this standard.
- Never add business logic to mocks.
- Always use `t.Cleanup` or `defer` for `AssertExpectations`.
- Place all mocks in test-only packages; never use in production code.

---

## 7. See Also

- [Handler Testing Standard](./handler_testing_standar.md)
- [Object Mother Guidelines](./object_mother_testing.md)

---

**By following this standard, all mocks in the project will be robust, maintainable, and easy for both humans and AI to generate and use.**
