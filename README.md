A fast and flexible CLI tool to generate Go project structures from embedded templates.  
Choose a template interactively and start building your app immediately.

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
