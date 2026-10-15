# Go Modular Monolith

A production-oriented **starter blueprint for a modular monolith** written in Go, paired with a
modern React web app — organised as a single monorepo and designed to grow into an ERP-style
system without splitting into microservices.

Each business capability lives in its own self-contained module with clear boundaries
(`api` → `app` → `domain` → `infra`). Modules are wired together at runtime through
dependency injection, so features can be added, enabled, or disabled without touching the rest
of the codebase.

## Highlights

- **Modular monolith** — one deployable binary, many independent domain modules.
- **Dependency injection** via [Uber Fx](https://uber-go.github.io/fx/); every module exposes an `fx.Module` that is composed in `cmd/server`.
- **CQRS & messaging** built on [Watermill](https://watermill.io/) with a durable PostgreSQL transport (or in-memory for dev), including a command/event bus, retries, and a poison queue.
- **Persistence** with PostgreSQL (`pgx` + `sqlx`), schema migrations via [Goose](https://github.com/pressly/goose), and type-safe query generation with [sqlc](https://sqlc.dev/) / [go-jet](https://github.com/go-jet/jet).
- **Authentication & authorization** — JWT auth, sessions in Redis, plus role-based (RBAC) and attribute-based (ABAC) access control with audit logging, API keys, and security events.
- **Multi-tenancy** primitives (single / shared / dedicated modes, `X-TENANT-ID` header).
- **Backend jobs** via a background job runner backed by [River](https://riverqueue.com/) and Redis.
- **File storage** abstraction with local and S3 backends.
- **Observability** — structured logging, metrics, and OpenTelemetry tracing hooks.
- **Typed, layered configuration** loaded from `configs/` with per-environment overrides.
- **OpenAPI-first frontend** — the Go server publishes an OpenAPI spec the web client is generated from.
- **Internationalisation** with [Paraglide](https://inlang.com/m/gerre34r/library-inlang-paraglideJs) across backend messages and UI strings (English/Arabic base).

## Tech Stack

**Backend**

| Area | Tools |
| --- | --- |
| Language | Go 1.27 |
| HTTP / API | [GoFrame v2](https://goframe.org/) (`ghttp`) with automatic OpenAPI |
| DI | Uber Fx |
| Config | [koanf](https://github.com/knadh/koanf) (YAML) |
| Database | PostgreSQL (`pgx`, `sqlx`), migrations via Goose |
| Code generation | sqlc, go-jet |
| Messaging | Watermill (`watermill-sql` PostgreSQL transport) |
| Jobs | River |
| Cache / sessions | Redis |
| Auth | `golang-jwt`, RBAC + ABAC |
| Storage | Local filesystem, AWS S3 |
| Observability | Zap, GoFrame `glog`, OpenTelemetry |

**Frontend**

| Area | Tools |
| --- | --- |
| Framework | [TanStack Start](https://tanstack.com/start) + React 19 |
| Routing / data | TanStack Router, TanStack Query, TanStack Table |
| Build | Vite 8, Bun workspaces |
| Styling / UI | Tailwind CSS v4, Base UI, Lucide icons |
| Forms / validation | React Hook Form + Zod |
| State | Zustand |
| API client | `openapi-fetch` / `openapi-react-query` (generated from the Go OpenAPI spec) |
| Quality | Biome |
| i18n | Paraglide |

## Repository Layout

```
.
├── cmd/server/            # Application entrypoint (composes modules + libs)
├── internal/              # Business modules (the heart of the app)
│   ├── auth/              # Auth module: api / app / domain / infra / migrations
│   └── adapters/          # Adapters that bridge modules
├── lib/                   # Reusable, module-agnostic libraries
│   ├── app/               # App bootstrap, HTTP server, middleware, OpenAPI
│   ├── auth/              # JWT, sessions, guard, RBAC/ABAC engine
│   ├── cqrs/              # Command bus, event bus, handler registry
│   ├── events/            # Watermill transport + message router
│   ├── db/                # Connection pool, transactions, repositories
│   ├── config/            # Typed configuration + YAML loading
│   ├── storage/           # File storage abstraction (local, S3)
│   ├── errors/ http/      # Shared error model and HTTP helpers
│   ├── jobs/ queues/      # Background jobs and queue consumers
│   ├── kv/ redis/         # Key-value and Redis clients
│   ├── tenancy/           # Tenant resolution & request context
│   └── logger/ utils/ ... # Logging and helpers
├── configs/               # base/*.yaml + environments/{dev,staging,production}
├── apps/web/              # TanStack Start web application
├── packages/ui/           # Shared UI component package
├── i18n/ messages/        # Localization catalogs
└── scripts/               # Localization + API client code generation
```

## The Auth Module (reference implementation)

`internal/auth` is the fully implemented module and the pattern every other module should follow:

```
internal/auth/
├── api/rest/   # HTTP controllers (authentication, account, users, rbac, abac, audit, apikeys)
├── app/        # Application services (use cases) + repository wiring
├── domain/     # Entities, value objects, repository interfaces, domain events
├── infra/      # sqlc-generated queries, repositories, ABAC resolver/cache
├── migrations/ # Goose SQL migrations
└── module.go   # fx.Module that wires repositories, services, and routes
```

Adding a new module means creating the same layout and registering its `fx.Module` in
`cmd/server/main.go`.

## Getting Started

### Prerequisites

- Go 1.27+
- [Bun](https://bun.sh/) (frontend workspace)
- Docker (for local PostgreSQL & Redis)

### 1. Start infrastructure

```bash
docker compose up -d
```

This launches PostgreSQL (port `5433`) and Redis (port `6379`).

### 2. Run the backend

```bash
go run ./cmd/server
```

The API listens on the port from `configs/environments/development.yaml` (default `8000`) and
exposes OpenAPI docs at `/api-docs` and the raw spec at `/api-docs-json`.

### 3. Run the web app

```bash
bun install
cd apps/web
bun --bun run dev
```

The web app runs on `http://localhost:3000`.

## Configuration

Configuration is split into composable YAML files:

- `configs/base/*.yaml` — defaults for app, auth, database, http, logging, messaging,
  observability, storage, integrations, jobs, modules, redis, queues, security, watermill, subscription.
- `configs/environments/*.yaml` — per-environment overrides (`development`, `staging`, `production`).

`lib/config` merges these into a single typed `Config` struct. Which feature modules are active
is controlled in `configs/base/modules.yaml`, and cross-cutting features (CQRS buses, queues,
jobs, integrations) can be toggled individually.

## Code Generation

Localization (Paraglide):

```bash
bun run i18n-gen
```

OpenAPI client for the web app (requires the backend to be running):

```bash
bun run ./scripts/generate-api-client.ts
```

Database query types are generated per module from the SQL in `internal/<module>/sql` using the
module's `sqlc.yaml`.

## Quick Start (GoFrame)

- https://goframe.org/quick
- https://goframe.org/docs/design/project-mono-repo
