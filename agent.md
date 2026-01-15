# AI Agent Instructions for photography-api

These instructions are tailored for AI assistants and contributors working on the `photography-api` project. Follow these guidelines to ensure all code aligns with the project's architecture, conventions, and quality standards.

---

## Project Overview

- **Language:** Go (Golang)
- **Architecture:** Domain-Driven Design (DDD) with Hexagonal (Ports & Adapters) structure
- **Core Location:** All main business logic and architecture reside under the [`/internal`](./internal) directory.

---

## Architectural Principles

- **Hexagonal Architecture:** Structure code into domain, application, and infrastructure layers. Use ports (interfaces) to decouple domain logic from external systems.
- **Domain-Driven Design:** Model the core business logic in the domain layer. Use value objects, entities, and domain services to represent business concepts.
- **Bounded Contexts:** Organize code into clear modules or packages representing distinct business areas.

---

## Directory & Layer Conventions

- **Domain Layer (`/internal/<context>/domain`):**
  - Contains entities, value objects, domain services, and business rules.
  - No dependencies on external packages or frameworks.
- **Application Layer (`/internal/<context>/application`):**
  - Contains use cases and application services.
  - Orchestrates domain logic and coordinates between domain and infrastructure.
- **Infrastructure Layer (`/internal/<context>/infrastructure`):**
  - Implements ports (interfaces) defined in the domain/application layers.
  - Handles persistence, external APIs, messaging, etc.
- **Adapters (`/internal/<context>/adapter`):**
  - Entry points (HTTP handlers, CLI, etc.) and secondary adapters (DB, external services).
  - Translate between external representations and internal models.

---

## Coding Guidelines

- **Follow Existing Patterns:** Study the structure and style in `/internal` and match new code to these conventions.
- **Interfaces for Ports:** Define interfaces in the domain/application layers. Implement them in infrastructure/adapters.
- **Dependency Direction:** Dependencies must always point inward (infra → app → domain).
- **Testing:** 
  - Write unit tests for domain logic.
  - Use mocks/fakes for infrastructure in tests.
  - Cover use cases with application-level tests.
  - **Always consult and follow the [main testing documentation](./docs/testing.md) and referenced standards for handler and object mother testing.**
- **Documentation:** 
  - Document the purpose and usage of each package, interface, and complex function.
  - Add comments for non-obvious business logic or architectural decisions.
- **Error Handling:** 
  - Return descriptive errors.
  - Use Go idioms for error wrapping and propagation.
- **Security:** 
  - Validate and sanitize all external input.
  - Follow Go security best practices.

---

## Collaboration & Maintenance

- **Consistency:** Maintain code style and naming conventions found in the project.
- **Extensibility:** Write code that is easy to extend and adapt as business requirements evolve.
- **TODOs & Issues:** Clearly mark incomplete features or technical debt with TODO comments and reference related issues.
- **Pull Requests:** Reference relevant issues and describe architectural impact in PRs.

---

*For any new feature or refactor, always reference the structure and patterns in [`/internal`](./internal), the [main testing documentation](./docs/testing.md) (which links to handler and object mother standards), and the [`/docs`](./docs) folder for any additional or future documentation. When in doubt, consult the docs to improve your responses and code suggestions.*
