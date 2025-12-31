# E-commerce Backend API

A robust, scalable e-commerce backend API built with Go, Gin framework, PostgreSQL, and Redis. This project follows clean architecture principles with a layered structure for maintainability and scalability.

## 🚀 Features

- **Authentication & Authorization**
  - OTP-based authentication via mobile number
  - JWT token-based authentication with RSA signing
  - Token revocation and expiration management

- **User Management**
  - User profile management
  - User registration and authentication
  - Profile updates

- **API Features**
  - RESTful API design
  - Request validation with custom error messages
  - Rate limiting and throttling
  - Structured logging with Zap
  - Health check endpoints
  - Custom error handling and recovery

- **Database & Caching**
  - PostgreSQL database with GORM ORM
  - Redis caching layer
  - Database migrations
  - Connection pooling

- **Security**
  - RSA-based JWT tokens
  - Input validation
  - SQL injection protection (via GORM)
  - Secure password handling

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

## ⚙️ Configuration

The application uses environment variables for configuration. Create a `.env` file in the root directory:

```env
# Server Configuration
APP_SERVER_PORT=8080
APP_SERVER_MODE=debug
APP_SERVER_READ_TIMEOUT=10s
APP_SERVER_WRITE_TIMEOUT=10s

# Database Configuration
APP_DATABASE_HOST=localhost
APP_DATABASE_PORT=5432
APP_DATABASE_USER=user
APP_DATABASE_PASSWORD=pass
APP_DATABASE_NAME=ecommerce
APP_DATABASE_SSL_MODE=disable
APP_DATABASE_MAX_OPEN_CONNS=25
APP_DATABASE_MAX_IDLE_CONNS=10
APP_DATABASE_CONN_MAX_LIFETIME=60m
APP_DATABASE_LOG_LEVEL=info

# Redis Configuration
APP_REDIS_HOST=localhost
APP_REDIS_PORT=6379
APP_REDIS_PASSWORD=
APP_REDIS_DATABASE=0
APP_REDIS_POOL_SIZE=10

# JWT Configuration
APP_JWT_ACCESS_TOKEN_DURATION=15m
APP_JWT_REFRESH_TOKEN_DURATION=24h
APP_JWT_SECRET=your-secret-key
APP_JWT_REFRESH_SECRET=your-refresh-secret-key

# Application Configuration
APP_APP_ENVIRONMENT=development
APP_APP_LOG_LEVEL=debug

# OTP Configuration
APP_OTPCODE_EXPIRE_TIME=3m
APP_OTPCODE_TRY_ATTEMPT=3
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

## 📡 API Endpoints

### Health Check

- `GET /health` - Health check endpoint (checks database and Redis connectivity)

### Authentication

- `POST /api/v1/auth/get-verification-code` - Request OTP code
  ```json
  {
    "mobile": "09123456789"
  }
  ```

- `POST /api/v1/auth/verify` - Verify OTP and get access token
  ```json
  {
    "mobile": "09123456789",
    "code": "1234"
  }
  ```

- `GET /api/v1/auth/logout` - Logout (requires authentication)

### User Management

- `GET /api/v1/client/user/profile` - Get user profile (requires authentication)
- `PUT /api/v1/client/user` - Update user profile (requires authentication)
  ```json
  {
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com"
  }
  ```

### Categories

- `GET /api/v1/category` - Get categories list

## 📁 Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── config/
│   └── config.go                # Configuration management
├── internal/
│   ├── dto/                     # Data Transfer Objects
│   │   ├── auth/
│   │   └── client/
│   ├── enums/                   # Enumerations
│   ├── handler/                 # HTTP handlers
│   │   ├── auth/
│   │   ├── client/
│   │   └── health/
│   ├── middlewares/             # HTTP middlewares
│   ├── model/                   # Database models
│   ├── repository/              # Data access layer
│   ├── resources/                # Response resources
│   ├── routers/                 # Route definitions
│   └── service/                 # Business logic layer
│       ├── auth/
│       └── client/
├── pkg/                          # Shared packages
│   ├── cache/                   # Redis client
│   ├── converter/               # Type converters
│   ├── database/                # Database connection
│   ├── logging/                 # Logger setup
│   ├── response/                # HTTP response helpers
│   └── validation/              # Validation utilities
├── keys/                         # RSA keys for JWT
├── docker-compose.yml            # Docker Compose configuration
├── Dockerfile                    # Docker image definition
├── go.mod                        # Go module definition
└── README.md                     # This file
```

## 🏗️ Architecture

The project follows a clean architecture pattern with clear separation of concerns:

- **Handler Layer**: HTTP request/response handling
- **Service Layer**: Business logic
- **Repository Layer**: Data access
- **Model Layer**: Database models

### Request Flow

```
HTTP Request → Middleware → Handler → Service → Repository → Database
                                    ↓
                              Response ← Resource
```

## 🔒 Security

- **JWT Authentication**: RSA256 signed tokens
- **Input Validation**: Comprehensive request validation
- **Rate Limiting**: Request throttling to prevent abuse
- **Error Handling**: Secure error messages (no sensitive data exposure)
- **SQL Injection Protection**: GORM parameterized queries

## 🧪 Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
go fmt ./...
```

### Building

```bash
go build -o bin/api cmd/api/main.go
```

## 🐳 Docker

### Build Docker Image

```bash
docker build -t ecommerce-backend .
```

### Run Container

```bash
docker run -p 8080:8080 --env-file .env ecommerce-backend
```

## 📝 Environment Variables

All configuration is done through environment variables with the `APP_` prefix. The application supports:

- `.env` file
- Environment variables
- Default values (for development)

## 🔍 Health Checks

The application provides health check endpoints for monitoring:

- **Health Check**: `GET /health` - Returns overall system health
- Checks database connectivity
- Checks Redis connectivity
- Returns detailed status for each service

Example response:
```json
{
  "success": true,
  "message": "success",
  "data": {
    "status": "healthy",
    "timestamp": "2024-01-15T10:30:00Z",
    "services": {
      "database": {
        "status": "healthy",
        "message": "Database connection is active",
        "response_time": "2.5ms"
      },
      "redis": {
        "status": "healthy",
        "message": "Redis connection is active",
        "response_time": "1.2ms"
      }
    }
  }
}
```

## 🚨 Error Handling

The API returns standardized error responses:

```json
{
  "success": false,
  "error": "Error message",
  "data": null
}
```

For validation errors:
```json
{
  "success": false,
  "error": "Validation failed",
  "data": {
    "errors": [
      {
        "field": "email",
        "message": "Must have a valid email address"
      }
    ]
  }
}
```

## 📚 API Response Format

All API responses follow a consistent format:

**Success Response:**
```json
{
  "success": true,
  "message": "success",
  "data": { ... }
}
```

**Error Response:**
```json
{
  "success": false,
  "error": "Error message",
  "data": null
}
```

## 🔐 Authentication

The API uses JWT tokens for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-token>
```

## 📊 Database Migrations

Database migrations run automatically on application startup. The application uses GORM's AutoMigrate feature.

## 🛠️ Troubleshooting

### Database Connection Issues

- Verify PostgreSQL is running
- Check database credentials in `.env`
- Ensure database exists

### Redis Connection Issues

- Verify Redis is running
- Check Redis configuration in `.env`

### JWT Key Issues

- Ensure `keys/private.pem` and `keys/public.pem` exist
- Verify key permissions

## 📄 License

[Add your license here]

## 👥 Contributors

[Add contributors here]

## 📞 Support

For issues and questions, please open an issue in the repository.

