# AGENTS.md

## Purpose

This repository contains a modular monolith platform implemented in Go.

The system follows a domain-oriented modular architecture where each business module owns its domain logic, application services, database access, HTTP endpoints, events, and integrations.

Agents operating on this repository must follow all specifications in this document.

---

# Technology Stack

| Category             | Technology                          |
| -------------------- | ----------------------------------- |
| Language             | Go                                  |
| Runtime              | Go 1.25+                            |
| Architecture         | Modular Monolith                    |
| Dependency Injection | Uber Fx                             |
| HTTP Framework       | GoFrame Routing                     |
| Database             | PostgreSQL                          |
| SQL Access           | sqlx + jet genrator & query builder |
| Migrations           | Goose                               |
| Cache                | Redis                               |
| Session Storage      | Redis                               |
| Messaging            | Watermill                           |
| CQRS                 | Watermill CQRS                      |
| Logging              | glog                                |
| Configuration        | Environment Variables               |
| Validation           | gvalid                              |
| Testing              | Go Testing Package                  |
| Containerization     | Docker                              |
| API Documentation    | OpenAPI                             |
| Authorization        | Session-Based Authentication        |
| Serialization        | JSON                                |

---

# Architectural Principles

## 1. Modular Monolith

The application is a single deployable unit.

Business functionality is organized into independent modules.

Modules communicate through:

- Commands
- Queries
- Domain Events
- Public Module APIs

Direct access to another module's database layer is prohibited.

---

## 2. Module Ownership

Each module owns:

- Domain models
- Repositories
- Application services
- CQRS handlers
- HTTP controllers
- Events
- Database schema

No module may directly modify another module's data.

---

## 3. Dependency Direction

Dependencies must flow inward.

Allowed:

Infrastructure → Application → Domain

Forbidden:

Domain → Application

Domain → Infrastructure

Application → HTTP Layer

---

## 4. CQRS

Commands mutate state.

Queries read state.

Commands must never return entities.

Queries must never modify state.

---

## 5. Event Driven Communication

Cross-module communication should prefer events.

Examples:

- UserCreated
- OrderApproved
- InvoicePaid
- StockReserved

Watermill is the standard event bus.

---

# Repository Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── users/
│   ├── auth/
│   ├── customers/
│   ├── products/
├──lib
│ ├── app
│ ├── auth
│ ├── common
│ ├── cache
│ ├── config
│ ├── cqrs
│ ├── db
│ ├── di
│ ├── errors
│ ├── events
│ ├── queues
│ ├── redis
│ └── tenancy
├── deployments/
├── docs/
├── configs/
├── tools/
└── tests/
```

---

# Module Structure

Every module must follow the same structure.

```text
module/

├── module.go
│
├── domain/
│   └── domain-name/
│       ├── entity.go
│       ├── value_objects.go
│       ├── events.go
│       └── repository.go

├── app/
│   ├── command/
│   ├── query/
│   └── app.go

├── infra/
│
│   └── entity-name/
│       └── entity-name.go
├── api/
│   ├── rest/
│   │   ├── controller-name.go
│   │   └── rest.go
│   │
│   └── middleware/
│
├── migrations/
└── i18n/
    ├── en.yaml
    └── ar.yaml
```

---

# Dependency Injection

All dependencies must be registered using Uber Fx.

Each module exposes:

```go
var Module = fx.Module(
    "users",
    ...
)
```

Modules must not instantiate dependencies directly.

Forbidden:

```go
db := sqlx.Open(...)
logger := glog.New(...)
```

Use dependency injection.

---

# Database Standards

## Database

PostgreSQL is the source of truth.

## Query Layer

sqlx & jet is the standard SQL access layer.

Raw SQL is allowed only when:

> [!NOTE]
> always include the schema name in raw queries because every module has its own db schema

- Performance requires it
- Query complexity requires it

## Migrations

Goose is the only migration tool.

Migration naming:

```text
202606200001_create_users_table.sql
202606200002_create_roles_table.sql
```

Creating migration:

Use the cli:

```sh
go run tools/cli/main.go migration:create name --module module-name

```

## go-jet

Use the cli to run migration and generate go jet types and structs :

```sh
go run tools/cli/main.go db:migrate -m module-name

```

---

# Repository Pattern

Repositories belong to the domain.

Interfaces:

```go
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    FindByID(ctx context.Context, id uuid.UUID) (*User, error)
}
```

Implementations belong to infra.

Common rules :

- Use soft-delete when possible
- Update the updatedAt fields when updating any record
- Avoid common mistakes such as sql injection

---

# Cache Standards

Redis is the standard cache.

Allowed uses:

- Sessions
- Query caching
- Rate limiting
- Distributed locks
- Temporary state

Redis must never be the source of truth.

---

# Authentication

Authentication is implemented as a dedicated module.

Responsibilities:

- Login
- Logout
- Session Management
- Password Verification
- Permission Resolution

Authentication middleware is exposed through public APIs.

Other modules consume middleware through dependency injection.

---

# HTTP Layer

Controllers must be thin.
Controllers contains:

- Request type struct
- Response type struct
- Controller struct with handlers methods
  like:

```go
type CreateTaskRequest struct {
 g.Meta              `path:"/" method:"post" tags:"Tasks" summary:"Create a task"`
 Name      string    `json:"name" v:"required" `
 DueDate   time.Date `json:"dueDate" v:"required"`
}
type CreateTaskResponse struct {
 g.Meta        `mime:"application/json"`
 Id     string `json:"id"`
 Name   string `json:"name"`
}
func (s *ModuleController) CreateTask(ctx context.Context, req *CreateTaskRequest) ( http.Response[CreateTaskResponse],  error) {
 // handle Create request
// if error
  if err!=nil{
    s.Fail(ctx,err)
  }
 return s.OK(data), nil
}

```

Controllers may:

- Validate requests
- Execute commands
- Execute queries
- Return responses

Controllers must not:

- Execute SQL
- Implement business rules
- Access repositories directly

All controllers are registered in the rest.go file

```go
// Package rest
// registers all controllers and handlers
// to the http server
package rest

import (
 "modular-monolith/lib/di"

 "github.com/gogf/gf/v2/net/ghttp"
)

func RegisterRoutes(in di.AppBaseParams,app *app.Application) {
 srv := in.Server
 srv.Group("/tasks", func(group *ghttp.RouterGroup) {
  group.Bind(NewModuleController(in,app))
 })
}
type ModuleController struct {
 di.AppBaseController
 sessionManager *auth.JWTManager
 App            *app.Application
}

func NewModuleController(deps di.AppBaseParams, app *app.Application) *ModuleController {
 return &ModuleController{
  AppBaseController: di.NewBaseController(deps),
  sessionManager:    auth.NewJWTManager(deps.Redis, deps.Config.Jwt),
  App:               app,
 }
}

```

---

# Validation

All request validation must use:

```go
github.com/gogf/gf/v2/util/gvalid
```

Validation belongs to request DTOs.

---

# Logging

glog is the standard logger.

Rules:

- Structured logging only
- No fmt.Println
- No log.Println

Required fields:

- request_id
- user_id
- module
- operation

---

# Error Handling

Errors must be typed.

Example:

```go
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidPassword = errors.New("invalid password")
)
```

No string comparisons.

---

# Configuration

Configuration comes from environment variables.

Access configuration through:

```go
lib/config
```

Do not read environment variables directly outside configuration packages.

Forbidden:

```go
os.Getenv("DATABASE_URL")
```

---

# Testing Standards

Required:

- Unit Tests
- Integration Tests

Preferred:

- Table Driven Tests

Example:

```go
func TestCreateUser(t *testing.T) {
    tests := []struct{
        name string
    }{
        ...
    }
}
```

---

# Naming Conventions

## Packages

```text
users
inventory
sales
accounting
```

Use lowercase.

---

## Files

```text
controller.go
service.go
repository.go
handler.go
module.go
```

Use snake_case only when necessary.

---

## Interfaces

```go
type UserRepository interface {}
type PasswordHasher interface {}
```

Do not prefix interfaces with I.

---

# Prohibited Patterns

Do not introduce:

- Service Locator
- Global State
- Circular Dependencies
- Active Record
- God Services
- Fat Controllers
- Direct Module Database Access
- Shared Mutable State

---

# Agent Rules

When modifying code:

1. Preserve module boundaries.
2. Prefer service methods over CQRS handlers .
3. Prefer events for cross-module communication.
4. Keep controllers thin.
5. Use dependency injection.
6. Follow repository interfaces.
7. Keep domain logic inside domain/application layers.
8. Never bypass module APIs.
9. Do not introduce framework coupling into domain models.
10. Maintain compile-time safety whenever possible.
11. Always wrap errors
12. When creating helper functions if the function would be used or could be use in multiple places add it to lib/utils package

---

# Git Guidelines

### General Rules

- Never rewrite Git history unless explicitly instructed.
- Never force push (`git push --force` or `git push --force-with-lease`) unless explicitly requested.
- Never delete branches, tags, or commits unless explicitly requested.
- Never modify Git configuration unless instructed.
- Never create commits automatically. Only commit when the user explicitly asks.
- Never create tags or releases unless requested.
- Never stash changes unless requested.
- Never discard uncommitted changes without explicit confirmation.

### Before Making Changes

- Check the repository status using:

  ```bash
  git status --short
  ```

- If the working tree contains unrelated modifications, do not overwrite them.
- Only modify files required for the current task.

### Branches

- Work on the currently checked-out branch.
- Do not switch branches automatically.
- Do not create new branches unless explicitly instructed.

### Commits

When asked to commit:

- Review the staged changes.
- Write concise commit messages following Conventional Commits.

Examples:

```
feat(auth): add JWT refresh endpoint
fix(users): prevent duplicate email creation
refactor(db): simplify transaction handling
docs(api): update authentication examples
test(cache): add redis integration tests
chore(deps): update Go dependencies
```

### Pulling

Before pulling:

- Ensure there are no uncommitted conflicting changes.
- Prefer:

```bash
git pull --rebase
```

unless the repository workflow specifies merge commits.

### Pushing

- Push only the current branch.
- Never push automatically after committing.
- Wait for explicit user approval before pushing.

### Generated Files

- Do not commit generated files unless they are intentionally version-controlled by the repository.
- Respect `.gitignore`.

### Large Refactors

For changes affecting many files:

1. Explain the planned scope.
2. Complete the refactor.
3. Run formatting and tests if available.
4. Present a summary before committing.

### Safety

Before any destructive Git operation, require explicit user confirmation.

Examples include:

- reset
- restore
- clean
- rebase
- cherry-pick
- revert
- branch deletion
- tag deletion
- force push
- history rewriting

### Commit Quality Checklist

Before creating a commit:

- Code builds successfully.
- Tests pass (if available).
- Formatting has been applied.
- No debug code remains.
- No commented-out code is introduced.
- No secrets, credentials, or API keys are committed.
- Commit contains only changes related to a single logical task.

Failure to follow these rules should be considered a design violation.
