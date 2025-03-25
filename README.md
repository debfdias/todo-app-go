# Todo API Server (Go Backend) 🚀

*Frontend integration (Next.js) in progress - coming soon!*

## Table of Contents
- [Features](#features)
- [API Documentation](#api-documentation)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Development](#development)
- [Deployment](#deployment)
- [Project Structure](#project-structure)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## Features
✔ RESTful API with Gin framework  
✔ PostgreSQL database with GORM ORM  
✔ JWT Authentication (coming in v1.1)  
✔ Automated database migrations  
✔ Health check endpoint with monitoring  
🚧 Next.js frontend integration in progress  

## API Documentation

### Base URL
`http://localhost:8080/`

### Endpoints
| Endpoint         | Method | Description                          |
|------------------|--------|--------------------------------------|
| `/health`        | GET    | Service status check (🚀 with emoji) |
| `/todos`         | GET    | List all todos                       |
| `/todos`         | POST   | Create new todo                      |
| `/todos/:id`     | GET    | Get single todo                      |
| `/todos/:id`     | PUT    | Update todo                          |
| `/todos/:id`     | DELETE | Remove todo                          |

## Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL 13+
- Docker (recommended)

### Installation
1. Clone repository:
   ```bash
   git clone https://github.com/yourrepo/todo-app.git
   cd todo-app/server
2. Setup env file:
   ```bash
   cp .env.example .env
3. Start database:
   ```bash
   gdocker-compose up -d
4. Start database:
   ```bash
   go mod download
5. Running locally:
   ```bash
   # With hot-reload (using Air)
    go install github.com/cosmtrek/air@latest
    air

    go run main.go
