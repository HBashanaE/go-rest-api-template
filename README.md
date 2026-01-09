# Go REST API Template

![Go Version](https://img.shields.io/github/go-mod/go-version/yourusername/go-api-template)
![CI Status](https://github.com/yourusername/go-api-template/actions/workflows/ci.yml/badge.svg)
[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy)

A scalable, production-ready Go REST API starter kit. Includes Docker, database integration, and CI/CD workflows.

## 🚀 Features

- **Structure**: Follows Standard Go Project Layout.
- **Routing**: Native Go 1.22+ `net/http` ServeMux (No heavy frameworks).
- **Database**: `pgx` driver for PostgreSQL.
- **Configuration**: Environment variable management via `godotenv`.
- **Deployment**: Docker, Docker Compose, and Render.com blueprint included.

## 🛠️ Tech Stack

- **Go**: 1.22+
- **PostgreSQL**: Database
- **Docker**: Containerization

## 🏁 Getting Started

### Prerequisites

- Go 1.22+
- Docker & Docker Compose

### Fast Spin-up (Docker)

```bash
# 1. Clone the repo
git clone [https://github.com/yourusername/go-api-template.git](https://github.com/yourusername/go-api-template.git)
cd go-api-template

# 2. Run everything
docker-compose up --build
```

The API will be available at http://localhost:8080/health.

## Local Development

1. Install dependencies
```
go mod download
```

2. Setup Environment
```
cp .env.example .env
# Update .env with your local DB credentials
```

3. Run
```
make run
```

4. Test
```
make test
```

## 📂 Project Structure

```
├── cmd/api/         # Main entry point
├── internal/        # Private application logic
│   ├── server/      # HTTP handlers and routing
│   └── database/    # Database connections
├── .github/         # CI/CD Workflows
└── docker-compose.yml
```