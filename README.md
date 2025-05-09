# BRB Mid-Svc Platform

A microservice written in Go for managing service bookings, vendors, and users. Features vendor summaries, booking logic, and direct PostgreSQL integration without an ORM.

## Features

* Vendor CRUD operations
* User management
* Service listing and management
* Booking creation with validation
* Vendor booking summary endpoint (`/summary/vendor/{id}`)
* Health check endpoint (`/health`)
* Swagger API documentation

## Tech Stack

* **Language:** Go (Golang)
* **Database:** PostgreSQL (using `database/sql` + `github.com/lib/pq`)
* **HTTP Server:** `net/http`
* **Documentation:** Swagger (via Swaggo)
* **Dependencies:** Git, Docker (optional)

## Project Structure

```
cmd/
  server/
    main.go           # application entrypoint
api/
  handler/
internal/
  repository/         # DB repositories
  usecase/            # business logic
  domain/             # domain models
docs/                 # generated Swagger docs
docker-compose.yml    # optional local setup
README.md             # this file
```

## Prerequisites

* Go 1.19+
* PostgreSQL
* `swag` CLI tool: `go install github.com/swaggo/swag/cmd/swag@latest`
* `http-swagger` package: `go get -u github.com/swaggo/http-swagger`

## Setup on a New Machine

1. **Clone the repository**

   ```bash
   git clone https://github.com/wikasdude/brb-midsvc-platform.git
   cd brb-midsvc-platform
   ```

2. **Install Go dependencies**

   ```bash
   go mod download
   ```

3. **Set environment variables**
   Create a `.env` file in the project root:

   ```dotenv
   POSTGRES_DSN=postgres://booking:booking@localhost:5432/brb?sslmode=disable
   ```

4. **Start local services** (optional via Docker Compose)

   ```bash
   docker-compose up -d
   ```

   This brings up PostgreSQL on `localhost:5432` with DB `brb`, user `booking`, and password `booking`.

5. **Generate Swagger docs**
   Ensure `cmd/server/main.go` has top-level annotations:

   ```go
   // @title BRB Mid-Svc Platform API
   // @version 1.0
   // @description API documentation for BRB Mid-Svc Platform
   // @host localhost:8080
   // @BasePath /
   ```

   Then run:

   ```bash
   swag init -g cmd/server/main.go
   ```

   Generated docs appear under `docs/`.

6. **Run the service**

   ```bash
   go run cmd/server/main.go
   ```

7. **Explore the API**

   * Health Check:  `GET http://localhost:8080/health`
   * Vendor Summary: `GET http://localhost:8080/summary/vendor/{vendor_id}`
   * Create Booking: `POST http://localhost:8080/bookings`
   * Swagger UI: `http://localhost:8080/swagger/index.html`
 
