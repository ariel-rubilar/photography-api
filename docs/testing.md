# Testing Standards – Main Guide

Welcome to the main testing documentation for the `photography-api` project. This guide provides an overview of our testing philosophy, structure, and references to detailed standards for specific testing patterns. All contributors—human or AI—must follow these standards to ensure high-quality, maintainable, and consistent tests across the codebase.

---

## 1. Overview

Testing in this project is guided by two core standards:

- **Handler Testing Standard:** How to write robust, maintainable tests for HTTP handlers.
- **Object Mother Testing Standard:** How to create and use test data builders (Object Mothers) for domain objects.

These standards are designed to be clear, enforceable, and AI-friendly.

---

## 2. Where to Find the Standards

- [Handler Testing Standard](./handler_testing_standard.md)  
  _Defines the required structure, helpers, and best practices for all HTTP handler tests._

- [Object Mother Testing Standard](./object_mother_testing.md)  
  _Explains how to build, organize, and use Object Mothers for generating test data._

---

## 3. General Testing Philosophy

- **Isolation:** Tests should isolate the unit under test, using mocks or fakes for dependencies.
- **Readability:** Tests must be easy to read and understand.
- **Repeatability:** Tests must be deterministic and produce the same result every run.
- **Maintainability:** Use helpers and patterns to avoid duplication and ease future changes.
- **Verification:** Always verify that mocks are called as expected and that all assertions are meaningful.

---

## 4. Test File Organization

- Place tests next to the code they test (e.g., handler tests next to handler implementations).
- Place reusable mocks and Object Mothers in dedicated `/test` folders within each context.
- Never import test helpers or Object Mothers into production code.

---

## 5. How to Get Started

1. **Read the [Handler Testing Standard](./handler_testing_standard.md)**  
   Learn how to structure handler tests, use helpers, and enforce mock verification.

2. **Read the [Object Mother Testing Standard](./object_mother_testing.md)**  
   Understand how to create and use Object Mothers for generating valid, customizable test data.

3. **Follow the checklist in each standard**  
   Each standard includes a checklist to ensure your tests meet project requirements.

---

## 6. For AI Contributors

- Always follow the referenced standards.
- Use existing helpers and Object Mothers—never invent ad-hoc patterns.
- Ensure all generated tests are readable, maintainable, and pass the checklists in the standards.

---

## 7. Updating the Standards

If you find gaps or improvements, propose changes via pull request. All changes must be reviewed for clarity, enforceability, and compatibility with both human and AI workflows.

---

**By following these standards, we ensure our tests are reliable, maintainable, and easy for anyone (or anything) to extend.**
