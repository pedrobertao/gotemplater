A fast and flexible CLI tool to generate Go project structures from embedded templates.  
Choose a template interactively and start building your app immediately.
<img width="545" alt="Screenshot 2025-05-29 at 18 38 58" src="https://github.com/user-attachments/assets/4ebf89ae-f7c0-4909-b7e9-098f064f1b65" />

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
