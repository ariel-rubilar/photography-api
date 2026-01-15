# Object Mother Guidelines – Testing Best Practices

## Purpose

The Object Mother pattern is used to **create fully valid domain objects for tests**.
This document provides **all rules and examples** for writing Object Mothers in this project.
It is intended for **developers and AI tools** to produce consistent, safe, and maintainable test code.

---

## 1. General Principles

1. **Test-only**

   * Object Mothers must live in test-only packages.
   * Never import them in production code.

2. **Valid by default**

   * `Default<Entity>()` must always return a fully valid domain object.
   * Avoid returning nil, invalid states, or partial objects.

3. **Customizable**

   * Use functional options (`With<Field>()`) to override defaults.
   * Avoid creating specialized constructors like `PhotoWithID`.

4. **Readable**

   * Tests should read like a story using Object Mothers.
   * Avoid hidden or surprising behavior in defaults.

5. **Deterministic when necessary**

   * Random data is allowed for general tests.
   * Override fields explicitly for assertions that require deterministic values.

---

## 2. Folder Structure

Object Mothers must be **inside the context test folder**:

```
/internal/<context>
  /domain
  /application
  /test
    /<entity>mother
      photo_mother.go
      recipe_mother.go
```

**Rules:**

* Test code may import `internal/<context>/test/*`.
* Production code must never import from `/test`.
* Each context owns its Object Mothers.

---

## 3. Naming Conventions

| Element            | Convention                |
| ------------------ | ------------------------- |
| Package            | `<entity>mother`          |
| Default factory    | `Default<Entity>()`       |
| Builder            | `New<Entity>(options...)` |
| Options            | `With<Field>()`           |
| Collection builder | `New<Entity>List(amount)` |

**Examples:**

* Package: `photomother`
* Default factory: `DefaultPhoto()`
* Builder: `NewPhoto()`
* Option: `WithID()`, `WithTitle()`
* List: `NewPhotoList(10)`

---

## 4. Implementation Rules

### Default Factory

```go
func DefaultPhoto() *photo.Photo {
	return photo.FromPrimitives(photo.PhotoPrimitives{
		ID:     gofakeit.UUID(),
		Title:  gofakeit.Sentence(3),
		URL:    gofakeit.URL(),
		Recipe: recipemother.DefaultRecipePrimitives(),
	})
}
```

* Must always return a valid object.
* May use realistic fake data.
* May depend on other Object Mothers in the same context.

### Functional Options

```go
type PhotoOption func(photo.PhotoPrimitives) photo.PhotoPrimitives
```

Each option modifies **only one field**:

```go
func WithID(id string) PhotoOption {
	return func(p photo.PhotoPrimitives) photo.PhotoPrimitives {
		p.ID = id
		return p
	}
}
```

### Builder Function

```go
func NewPhoto(options ...PhotoOption) *photo.Photo {
	primitives := DefaultPhoto().ToPrimitives()
	for _, opt := range options {
		primitives = opt(primitives)
	}
	return photo.FromPrimitives(primitives)
}
```

* Start from `Default<Entity>()`.
* Apply options sequentially.
* Return a real domain entity.

### List Builder

```go
func NewPhotoList(amount int) []*photo.Photo {
	photos := make([]*photo.Photo, amount)
	for i := range amount {
		photos[i] = NewPhoto()
	}
	return photos
}
```

* Use for pagination, bulk, or collection-based tests.

---

## 5. How to Use in Tests

### Simple test

```go
photo := photomother.DefaultPhoto()
```

### Override fields

```go
photo := photomother.NewPhoto(
	photomother.WithID("photo-1"),
	photomother.WithTitle("My Photo"),
)
```

### Deterministic assertions

```go
photo := photomother.NewPhoto(
	photomother.WithID("fixed-id"),
)
```

---

## 6. What NOT to Do

* ❌ Place Object Mothers in production packages.
* ❌ Return invalid objects in defaults.
* ❌ Create multiple specialized constructors.
* ❌ Share Object Mothers across unrelated contexts.
* ❌ Use hidden random values for fields under test assertion.

---

## 7. Duplication vs Reuse

* Duplication in tests is **acceptable** when:

  * Contexts are independent
  * Domain meaning diverges
  * Readability is reduced by reuse
* Prefer local duplication over coupling contexts.

---

## 8. AI Guidelines

AI tools generating tests **must**:

* Use existing Object Mothers
* Follow naming conventions
* Prefer `New<Entity>(WithX())` over inline construction
* Never bypass domain invariants
* Never create ad-hoc test objects in tests

---

## 9. Summary Checklist

* [ ] Object Mother lives in `_test.go` or `/test` folder
* [ ] Defaults return fully valid objects
* [ ] Functional options allow field overrides
* [ ] Collections builders are available for bulk tests
* [ ] Tests are readable and deterministic where needed
* [ ] Production code never imports Object Mothers
* [ ] Duplication is acceptable between independent contexts

> Following these rules ensures maintainable, clear, and AI-friendly tests in all contexts of the project.

```
```
