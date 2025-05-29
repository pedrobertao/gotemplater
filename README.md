# 🛠️ gotemplater [██████░░░░░░] 40% completed

A fast and flexible CLI tool to generate Go project structures from embedded templates.  
Choose a template interactively and start building your app immediately.

---

## ✨ Features

- 📦 Interactive CLI: choose a template by number
- 🔐 Embedded templates using Go's `embed.FS` — no external files needed
- 📁 Prebuilt templates for:
  - Simple Web API
  - Web API + MongoDB
  - Web API + PostgreSQL
  - Web API + SQLite
- 🧪 Auto-generated structure with handlers, services, repositories, and tests
- 🐳 Optional Docker and docker-compose support

---

## 📦 Installation

```bash
go install -u github.com/pedrobertao/gotemplater
```

## 📦 How to use

```bash
gotemplater
- Choose name
- Choose template
Done !
```

## Examples

web-api

my-api/
├── cmd/
│ └── server/
│ └── main.go # Application entry point
├── internal/
│ ├── api/
│ │ ├── handlers/
│ │ │ ├── user.go # User route handlers
│ │ │ ├── auth.go # Auth route handlers
│ │ │ └── health.go # Health check handlers
│ │ ├── middleware/
│ │ │ ├── auth.go # Authentication middleware
│ │ │ ├── cors.go # CORS middleware
│ │ │ └── logging.go # Logging middleware
│ │ └── routes/
│ │ └── routes.go # Route definitions
│ ├── config/
│ │ └── config.go # Configuration management
│ ├── database/
│ │ ├── connection.go # Database connection
│ │ └── migrations/
│ │ └── 001_create_users.sql
│ ├── models/
│ │ ├── user.go # User model/struct
│ │ └── response.go # API response models
│ ├── repository/
│ │ ├── interfaces.go # Repository interfaces
│ │ ├── user_repo.go # User repository implementation
│ │ └── postgres/
│ │ └── user.go # PostgreSQL specific implementation
│ ├── service/
│ │ ├── user_service.go # Business logic layer
│ │ └── auth_service.go # Authentication service
│ └── utils/
│ ├── jwt.go # JWT utilities
│ ├── validator.go # Request validation
│ └── response.go # Response helpers
├── pkg/
│ └── logger/
│ └── logger.go # Shared logging package
├── scripts/
│ ├── build.sh # Build scripts
│ └── migrate.sh # Database migration scripts
├── tests/
│ ├── integration/
│ │ └── api_test.go # Integration tests
│ └── unit/
│ ├── handlers_test.go # Unit tests for handlers
│ └── services_test.go # Unit tests for services
├── docs/
│ ├── api.md # API documentation
│ └── swagger.yaml # OpenAPI/Swagger spec
├── deployments/
│ ├── docker/
│ │ └── Dockerfile # Docker configuration
│ └── k8s/
│ ├── deployment.yaml # Kubernetes deployment
│ └── service.yaml # Kubernetes service
├── .env.example # Environment variables example
├── .gitignore # Git ignore file
├── go.mod # Go module file
├── go.sum # Go dependencies checksum
├── Makefile # Build automation
└── README.md # Project documentation
