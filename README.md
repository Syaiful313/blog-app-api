# Blog App API

RESTful API untuk aplikasi blog yang dibangun dengan Go Fiber, GORM, dan PostgreSQL. API ini menyediakan fitur autentikasi, manajemen blog, dan sistem CRUD yang lengkap.

## 🚀 Fitur Utama

- **🔐 Authentication System** - JWT-based auth dengan register, login, forgot password, dan reset password
- **📝 Blog Management** - CRUD operations untuk blog posts dengan slug generation
- **👥 User Management** - Sistem pengelolaan pengguna dengan profil lengkap  
- **📧 Email Integration** - Email notifications untuk reset password
- **🛡️ Authorization** - Role-based access control untuk resource protection
- **🔍 Pagination & Search** - Dukungan pagination dan pencarian untuk performa optimal
- **🐳 Docker Ready** - Containerization dengan Docker Compose
- **🔄 Hot Reload** - Development dengan Air untuk produktivitas tinggi
- **📋 Advanced Middleware** - CORS, Auth, Error handling yang robust
- **🧪 Testing Ready** - Struktur yang mendukung unit dan integration testing

## 📁 Struktur Proyek

```
blog-app-api/
├── cmd/                           # Entry point aplikasi
│   └── main.go
├── config/                        # Konfigurasi aplikasi
│   └── config.go
├── database/                      # Database connection & migration
│   └── database.go
├── internal/                      # Business logic (private)
│   ├── controllers/              # HTTP request handlers
│   │   ├── auth_controller.go
│   │   ├── blog_controller.go
│   │   └── sample_controller.go
│   ├── middleware/               # Custom middlewares
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── error.go
│   ├── models/                   # Data models & DTOs
│   │   ├── user.go
│   │   ├── blog.go
│   │   └── sample.go
│   ├── routes/                   # Route definitions
│   │   ├── routes.go
│   │   ├── auth_router.go
│   │   ├── blog_router.go
│   │   └── sample_router.go
│   └── services/                 # Business logic layer
│       ├── auth_service.go
│       ├── blog_service.go
│       └── sample_service.go
├── pkg/                          # Public packages
│   └── response/                # Standardized API responses
│       └── response.go
├── utils/                        # Utility functions
│   ├── jwt.go                   # JWT token management
│   ├── password.go              # Password hashing
│   ├── email.go                 # Email utilities
│   └── token.go                 # Token generation
├── .air.toml                    # Air configuration
├── .env.example                 # Environment template
├── docker-compose.yml           # Docker services
├── Makefile                     # Build automation
├── go.mod                       # Go modules
└── README.md
```

## 🛠️ Prerequisites

- **Go** 1.23+ 
- **Docker** & **Docker Compose**
- **Make** (optional, untuk build automation)
- **Air** (optional, untuk hot reload development)

## ⚡ Quick Start

### 1. Clone Repository
```bash
git clone https://github.com/Syaiful313/blog-app-api.git
cd blog-app-api
```

### 2. Environment Setup
```bash
cp .env.example .env
# Edit .env sesuai konfigurasi Anda
```

### 3. Install Dependencies
```bash
go mod download
go mod tidy
```

### 4. Start Services
```bash
# Start database & adminer
docker-compose up -d

# Wait for database to be ready
make docker-setup
```

### 5. Run Application
```bash
# Production mode
make start

# Development mode (dengan hot reload)
make dev
```

Server akan berjalan di `http://localhost:8000`

## ⚙️ Environment Configuration

Konfigurasi lengkap pada file `.env`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=go_fiber_db

# Server Configuration
PORT=8000
JWT_SECRET=your_jwt_secret_key_here

# CORS Configuration
CORS_ALLOWED_ORIGINS=*
CORS_ALLOW_CREDENTIALS=false

# Email Configuration (SMTP)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password
FROM_EMAIL=noreply@yourapp.com

# Frontend URL (untuk reset password links)
FRONTEND_URL=http://localhost:3000
```

## 🔗 API Endpoints

### Base URL
```
http://localhost:8000
```

### Health Check
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/health` | Server health status |

### Authentication
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/auth/register` | Register new user | ❌ |
| POST | `/auth/login` | User login | ❌ |
| POST | `/auth/forgot-password` | Request password reset | ❌ |
| POST | `/auth/reset-password` | Reset password with token | ❌ |

### Blog Management
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/blogs` | Get all blogs (paginated) | ❌ |
| GET | `/blogs/:id` | Get blog by ID | ❌ |
| POST | `/blogs` | Create new blog | ✅ |

### Samples (Demo CRUD)
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/samples` | Get all samples (paginated) | ❌ |
| GET | `/samples/:id` | Get sample by ID | ❌ |
| POST | `/samples` | Create new sample | ✅ |
| PUT | `/samples/:id` | Update sample | ✅ |
| DELETE | `/samples/:id` | Delete sample | ✅ |

## 📋 Request/Response Examples

### Register User
```bash
curl -X POST http://localhost:8000/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

**Response:**
```json
{
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "first_name": "John",
      "last_name": "Doe",
      "is_active": true,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    }
  }
}
```

### Login User
```bash
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "email": "user@example.com",
      "first_name": "John",
      "last_name": "Doe",
      "is_active": true,
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    }
  }
}
```

### Create Blog (Protected)
```bash
curl -X POST http://localhost:8000/blogs \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "My First Blog Post",
    "content": "This is the content of my first blog post...",
    "published": true
  }'
```

### Forgot Password
```bash
curl -X POST http://localhost:8000/auth/forgot-password \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com"
  }'
```

### Get Blogs with Pagination
```bash
curl "http://localhost:8000/blogs?page=1&limit=10"
```

## 🔥 Development

### Available Make Commands
```bash
make build          # Build aplikasi
make run            # Run aplikasi production
make dev            # Run dengan hot reload (Air)
make test           # Run semua tests
make clean          # Clean build artifacts
make docker-up      # Start Docker containers
make docker-down    # Stop Docker containers
make docker-setup   # Start containers + wait for DB
make deps           # Install dependencies
make dev-deps       # Install development tools
make fmt            # Format kode
make lint           # Lint kode (requires golangci-lint)
make setup          # Full development setup
```

### Hot Reload Development
```bash
# Install Air (jika belum ada)
make dev-deps

# Start development server
make dev
```

Air akan secara otomatis restart server ketika ada perubahan file Go.

### Database Management
```bash
# Start database
make docker-up

# Access database via Adminer
open http://localhost:8080

# Database credentials:
# System: PostgreSQL
# Server: postgres
# Username: postgres  
# Password: admin
# Database: go_fiber_db
```

## 🧪 Testing

```bash
# Run semua tests
make test

# Run tests dengan coverage
go test -v -cover ./...

# Run specific test
go test -v ./internal/services/
```

## 🐳 Docker Services

Docker Compose menyediakan:

### PostgreSQL Database
- **Port**: 5432
- **Username**: postgres
- **Password**: admin
- **Database**: go_fiber_db

### PostgreSQL Test Database  
- **Port**: 5433
- **Username**: postgres
- **Password**: admin
- **Database**: go_fiber_test_db

### Adminer (Database UI)
- **URL**: http://localhost:8080
- **Features**: Browse, edit, dan manage database via web interface

```bash
# Start semua services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Remove volumes (reset database)
docker-compose down -v
```

## 📧 Email Features

API mendukung email notifications untuk:

- **Password Reset**: Email dengan secure reset link
- **Reset Confirmation**: Konfirmasi setelah password berhasil direset

Template email menggunakan HTML responsive dengan styling modern.

## 🛡️ Security Features

- **JWT Authentication**: Secure token-based authentication
- **Password Hashing**: bcrypt untuk password security
- **CORS Protection**: Configurable CORS policies
- **Input Validation**: Request validation & sanitization
- **Error Handling**: Structured error responses tanpa sensitive data
- **Token Expiry**: Automatic token expiration (24 jam untuk auth, 1 jam untuk reset)

## 🚀 Deployment

### Binary Deployment
```bash
# Build binary
make build

# Run binary
./bin/main
```

### Docker Deployment
```bash
# Build production image
docker build -t blog-app-api .

# Run with environment variables
docker run -p 8000:8000 --env-file .env blog-app-api
```

## 📊 Performance Features

- **Pagination**: Efficient data loading dengan page & limit
- **Database Indexes**: Optimized query performance  
- **Connection Pooling**: GORM connection management
- **Middleware Stack**: Efficient request processing
- **Slug Generation**: SEO-friendly URLs untuk blog posts

## 🤝 Contributing

1. Fork repository ini
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit perubahan (`git commit -m 'Add amazing feature'`)
4. Push ke branch (`git push origin feature/amazing-feature`)
5. Buka Pull Request

### Development Guidelines
- Follow Go conventions dan best practices
- Write tests untuk fitur baru
- Update documentation jika diperlukan
- Ensure code passes linting (`make lint`)

## 📄 License

Distributed under the MIT License. Lihat file `LICENSE` untuk informasi lebih lengkap.

## 📧 Contact & Support

**Developer**: Syaiful  
**GitHub**: [@Syaiful313](https://github.com/Syaiful313)  
**Project Repository**: [blog-app-api](https://github.com/Syaiful313/blog-app-api)

---

⭐ **Star repository ini jika membantu Anda!**

## 🔗 Related Links

- [Go Fiber Documentation](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [JWT Best Practices](https://tools.ietf.org/html/rfc7519)