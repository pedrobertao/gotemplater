# 🛠️ gotemplater

A minimal and extensible project scaffolding tool for Go developers.  
Generate ready-to-code project structures with Docker, database configs, and tests — in one command.

---

## ✨ Features

- ✅ Template-based project generation via YAML
- 📁 Predefined structures for:
  - Basic Web Projects
  - PostgreSQL with Docker Compose
  - SQLite Repositories
- 🧪 Includes boilerplate for services, handlers, middleware, and tests
- 🐳 Docker and `docker-compose.yml` generation support
- 💨 Lightweight and installable via `go install`

---

## 📦 Installation

```bash
go install github.com/pedrobertao/gotemplater@latest
```

## 🚀 Usage

```bash
gotemplater -name=myapp -template=api-simple
```
