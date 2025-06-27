# passIt

A modern password management and sharing platform built with Go, Keycloak, and PostgreSQL.

## Table of Contents
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Database Migration](#database-migration)
- [Local Development](#local-development)
- [Makefile Commands](#makefile-commands)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [License](#license)

## Getting Started

Follow these instructions to set up the project for local development and testing.

## Prerequisites
- Go 1.24.0 or higher
- Docker & Docker Compose

## Database Migration

Run database migrations:
```bash
go run ./internal/database/migration/migration.go
```

## Local Developement

### Partial developement

With docker-compose.yml file a functional application will start
Start Core Services:
- Database
- Keycloack
```bash
docker-compose up -d
```
Run the Go application:
```bash
make run
```

### Full developement

Start all Services:
- Database
- Keycloack
- elasticsearch
- kibana(elasticSearch UI)
- filebeat
- Passit backend(If there is no need for that you can comment it out)
```bash
docker-compose -f docker-compose-full.yml up --build -d
```

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## Environment Variables
An .env file in the root directory is needed with the following variables:
```bash
PORT=8080
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=passit
DB_USERNAME=melkey
DB_PASSWORD=password1234
DB_SCHEMA=public
```
