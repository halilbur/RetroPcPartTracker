# Retro PC Part Tracker

Retro PC Part Tracker is a small server-rendered web application for browsing PC hardware parts and comparing dealer prices. It focuses on a simple catalog experience: recent parts, search, category filtering, and part detail pages with price and stock information.

The project is a Go application backed by PostgreSQL. HTML is rendered on the server with `templ`, routing is handled by Gin, and the initial database schema includes sample parts, dealers, and prices for local development.

## Architecture

```text
.
|-- main.go                 # Application entry point, DB connection, Gin routes
|-- handlers/               # HTTP handlers and template rendering
|-- store/                  # PostgreSQL-backed data access
|-- templates/              # templ views and generated Go template code
|-- static/css/             # Retro-styled CSS
|-- schema.sql              # Database schema, indexes, and sample data
|-- docker-compose.yml      # Local PostgreSQL and Redis services
|-- go.mod                  # Go module and dependencies
`-- main_test.go            # DATABASE_URL configuration tests
```

Request flow:

```text
Browser
  -> Gin route in main.go
  -> handler in handlers/
  -> PartStore query in store/
  -> PostgreSQL tables from schema.sql
  -> templ-rendered HTML response
```

Core routes:

| Route | Purpose |
| --- | --- |
| `GET /` | Shows search form and recently added parts. |
| `GET /search?q=...` | Searches parts by name, brand, or specs. |
| `GET /parts/:type` | Lists parts for a specific type, such as `GPU`, `CPU`, `RAM`, `Motherboard`, or `Storage`. |
| `GET /part/:id` | Shows part details and dealer price comparison. |

## Tech Stack

| Area | Technology | Notes |
| --- | --- | --- |
| Language | Go `1.25.5` | Declared in `go.mod`. |
| Web framework | Gin `v1.9.1` | HTTP routing and middleware. |
| Views | templ `v0.3.977` | Server-rendered HTML components. |
| Database | PostgreSQL 14 | Local service defined in Docker Compose. |
| Database driver | `github.com/lib/pq` `v1.10.9` | PostgreSQL driver for `database/sql`. |
| Styling | Plain CSS | Retro web aesthetic under `static/css/style.css`. |
| Local services | Docker Compose | Starts PostgreSQL and Redis. Redis is defined in compose but is not currently used by the Go application code. |

There is no `.csproj` in this repository; the application is organized as a Go module.

## Data Model

The PostgreSQL schema defines three main tables:

| Table | Purpose |
| --- | --- |
| `parts` | Hardware catalog entries with name, type, brand, specs, image URL, and creation time. |
| `dealers` | Dealer metadata including URL, authenticity rating, and verification flag. |
| `prices` | Dealer-specific part prices, currency, stock status, and last updated time. |

`schema.sql` also creates indexes for common lookup paths and inserts sample data for local development.

## Setup

### Prerequisites

- Go matching the version declared in `go.mod`
- Docker and Docker Compose
- A shell with environment variable support

### 1. Start local services

```bash
docker compose up -d
```

This starts:

- PostgreSQL on `localhost:5432`
- Redis on `localhost:6379`

PostgreSQL loads `schema.sql` through Docker's initialization mechanism when the database volume is first created.

### 2. Configure the database URL

The application requires `DATABASE_URL`.

```bash
export DATABASE_URL='postgres://postgres:postgres@localhost:5432/pcparts?sslmode=disable'
```

PowerShell equivalent:

```powershell
$env:DATABASE_URL = 'postgres://postgres:postgres@localhost:5432/pcparts?sslmode=disable'
```

### 3. Install dependencies

```bash
go mod download
```

### 4. Run the application

```bash
go run .
```

The server listens on:

```text
http://localhost:8080
```

### 5. Run tests

```bash
go test ./...
```

## Development Notes

- The generated templ output is committed as `templates/templates_templ.go`.
- If `templates/templates.templ` is changed, regenerate the templ output before running the app.
- The app reads database configuration from `DATABASE_URL`; no fallback connection string is embedded in code.
- Static assets are served from `/static`.

## Roadmap

Current project-facing next steps:

- Add a documented templ regeneration command or script.
- Add repeatable database migration workflow beyond the initial `schema.sql`.
- Add application-level containerization so the Go service can run alongside the existing Compose services.
- Expand automated tests around handlers and store queries.
- Clarify whether the Redis service should remain in the local stack or be wired into a concrete application use case.
