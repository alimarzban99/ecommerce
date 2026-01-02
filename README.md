# E-commerce Backend API

A robust, scalable e-commerce backend API built with Go, Gin framework, PostgreSQL, and Redis. This project follows clean architecture principles with a layered structure for maintainability and scalability.

## 📋 Prerequisites

- Go 1.25 or higher
- PostgreSQL 12 or higher
- Redis 7 or higher
- Docker and Docker Compose (optional, for containerized setup)

## 🛠️ Tech Stack

- **Framework**: Gin Web Framework
- **Language**: Go 1.25
- **Database**: PostgreSQL (with GORM)
- **Cache**: Redis
- **Authentication**: JWT (RSA256)
- **Validation**: go-playground/validator
- **Configuration**: Viper
- **Logging**: Zap Logger
- **ORM**: GORM

## 📦 Installation

### Clone the Repository

```bash
git clone <repository-url>
cd shop/backend
```

### Install Dependencies

```bash
go mod download
```

### JWT Keys Setup

You need to generate RSA key pairs for JWT signing. Place them in the `keys/` directory:

```bash
# Generate private key
openssl genrsa -out keys/private.pem 2048

# Generate public key
openssl rsa -in keys/private.pem -pubout -out keys/public.pem
```

## 🏃 Running the Application

### Local Development

1. **Start PostgreSQL and Redis** (using Docker Compose):

```bash
docker-compose up -d db redis
```

2. **Run the application**:

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

### Using Docker Compose

Run the entire stack with Docker Compose:

```bash
docker-compose up -d
```

This will start:
- Application server (port 8080)
- PostgreSQL database (port 5432)
- Redis (port 6379)
- pgAdmin (port 5050)


## 🏗️ Architecture

The project follows a clean architecture pattern with clear separation of concerns:

- **Handler Layer**: HTTP request/response handling
- **Service Layer**: Business logic
- **Repository Layer**: Data access
- **Model Layer**: Database models

