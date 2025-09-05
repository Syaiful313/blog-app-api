# Blog API - Go Fiber

RESTful API untuk aplikasi blog yang dibangun dengan Go Fiber, GORM, PostgreSQL, dan Cloudinary. API ini menyediakan sistem autentikasi JWT, manajemen blog dengan upload gambar, dan sistem CRUD yang lengkap.

## 🚀 Fitur Utama

- **🔐 Sistem Autentikasi** - JWT-based authentication dengan register, login, forgot password, dan reset password
- **📝 Manajemen Blog** - CRUD operations untuk blog posts dengan slug generation otomatis
- **🖼️ Upload Gambar** - Integrasi Cloudinary untuk upload dan manajemen gambar blog
- **📧 Email Notifications** - Email untuk reset password dengan template HTML yang responsive  
- **👥 User Management** - Sistem pengelolaan pengguna dengan profil lengkap
- **🛡️ Authorization** - Middleware auth untuk proteksi resource
- **🔍 Pagination** - Dukungan pagination untuk performa optimal
- **🐳 Docker Ready** - Containerization dengan Docker Compose
- **🔄 Hot Reload** - Development dengan Air untuk produktivitas tinggi

## 📁 Struktur Proyek

```
blog-app-api/
├── cmd/                          # Entry point aplikasi
│   └── main.go
├── config/                       # Konfigurasi aplikasi
│   └── config.go
├── database/                     # Database connection & migration
│   └── database.go
├── internal/                     # Business logic (private)
│   ├── controllers/             # HTTP request handlers
│   ├── middlewares/             # Custom middlewares
│   ├── models/                  # Data models & DTOs
│   ├── routes/                  # Route definitions
│   └── services/                # Business logic layer
├── pkg/                         # Public packages
│   └── response/               # Standardized API responses
├── utils/                       # Utility functions
│   ├── jwt.go                  # JWT token management
│   ├── password.go             # Password hashing
│   ├── email.go                # Email utilities
│   ├── token.go                # Token generation
│   └── cloudinary.go           # Cloudinary integration
├── .air.toml                   # Air configuration
├── .env.example                # Environment template
├── docker-compose.yml          # Docker services
├── Makefile                    # Build automation
└── README.md
```

## 🛠️ Prerequisites

- **Go** 1.23+
- **Docker** & **Docker Compose**
- **Cloudinary Account** (untuk upload gambar)
- **SMTP Email** (untuk fitur forgot password)

## ⚡ Quick Start

### 1. Clone Repository
```bash
git clone <repository-url>
cd blog-app-api
```

### 2. Environment Setup
```bash
cp .env.example .env
```

Edit file `.env` dengan konfigurasi Anda:
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=go_fiber_db

# Server Configuration
PORT=8000
JWT_SECRET=your_strong_jwt_secret_here

# Email Configuration (SMTP)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password
FROM_EMAIL=noreply@yourapp.com

# Frontend URL (untuk reset password links)
FRONTEND_URL=http://localhost:3000

# Cloudinary Configuration (Required untuk upload gambar)
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
```

### 3. Install Dependencies
```bash
go mod download
```

### 4. Start Database
```bash
docker-compose up -d
```

### 5. Run Application
```bash
# Production mode
make run

# Development mode (dengan hot reload)
make dev
```

Server akan berjalan di `http://localhost:8000`

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
| POST | `/auth/register` | Register pengguna baru | ❌ |
| POST | `/auth/login` | Login pengguna | ❌ |
| POST | `/auth/forgot-password` | Request reset password | ❌ |
| POST | `/auth/reset-password` | Reset password dengan token | ❌ |

### Blog Management
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/blogs` | Get semua blog (dengan pagination) | ❌ |
| GET | `/blogs/:id` | Get blog berdasarkan ID | ❌ |
| POST | `/blogs` | Create blog baru (dengan upload gambar) | ✅ |
| PATCH | `/blogs/:id` | Update blog (dengan upload gambar) | ✅ |
| DELETE | `/blogs/:id` | Delete blog | ✅ |

### Sample CRUD (Demo)
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/samples` | Get semua sample (dengan pagination) | ❌ |
| GET | `/samples/:id` | Get sample berdasarkan ID | ❌ |
| POST | `/samples` | Create sample baru | ✅ |
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

### Create Blog dengan Image Upload
```bash
curl -X POST http://localhost:8000/blogs \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "title=My First Blog Post" \
  -F "content=This is the content of my first blog post..." \
  -F "published=true" \
  -F "image=@/path/to/image.jpg"
```

**Response:**
```json
{
  "message": "Blog created successfully",
  "data": {
    "id": 1,
    "title": "My First Blog Post",
    "content": "This is the content of my first blog post...",
    "slug": "my-first-blog-post",
    "published": true,
    "image_url": "https://res.cloudinary.com/your-cloud/image/upload/...",
    "userId": 1,
    "user": {
      "id": 1,
      "email": "user@example.com",
      "first_name": "John",
      "last_name": "Doe"
    },
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z"
  }
}
```

### Get Blogs dengan Pagination
```bash
curl "http://localhost:8000/blogs?page=1&limit=10"
```

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "title": "My First Blog Post",
      "slug": "my-first-blog-post",
      "content": "Blog content...",
      "published": true,
      "image_url": "https://res.cloudinary.com/...",
      "userId": 1,
      "user": {
        "id": 1,
        "email": "user@example.com",
        "first_name": "John",
        "last_name": "Doe"
      },
      "created_at": "2025-01-15T10:30:00Z",
      "updated_at": "2025-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 25
  }
}
```

## 🔥 Development Commands

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
make setup          # Full development setup
```

### Hot Reload Development
```bash
# Install Air (development dependency)
go install github.com/cosmtrek/air@latest

# Start development server
make dev
```

## 🐳 Docker Services

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

### Adminer (Database Management UI)
- **URL**: http://localhost:8080
- **System**: PostgreSQL
- **Server**: postgres
- **Username**: postgres
- **Password**: admin
- **Database**: go_fiber_db

## 📧 Email Features

API mendukung email notifications dengan template HTML responsive untuk:

- **Password Reset**: Email dengan secure reset link (expires dalam 1 jam)
- **Reset Confirmation**: Konfirmasi setelah password berhasil direset

## 🛡️ Security Features

- **JWT Authentication**: Token-based auth dengan expiry 24 jam
- **Password Hashing**: bcrypt untuk keamanan password
- **CORS Protection**: Configurable CORS policies
- **Input Validation**: Request validation & sanitization
- **Authorization**: User-based resource access control
- **Secure File Upload**: Image validation dan size limiting
- **Token Expiry**: Automatic token expiration

## 📊 Features Unggulan

### Blog Management
- **Auto Slug Generation**: SEO-friendly URLs dari title
- **Image Upload**: Cloudinary integration dengan optimasi otomatis
- **User Authorization**: User hanya bisa edit/delete blog sendiri
- **Pagination**: Efficient data loading
- **Published Status**: Draft dan published state

### Image Handling
- **File Validation**: Type checking (JPEG, PNG) 
- **Size Limitation**: Maximum 2MB untuk blog images
- **Auto Optimization**: Cloudinary auto-format dan quality
- **Secure Storage**: Cloud-based dengan CDN

### Authentication System
- **JWT Tokens**: Stateless authentication
- **Password Reset Flow**: Email-based dengan secure tokens
- **User Sessions**: 24-hour token validity
- **Registration**: Email validation dan password requirements

## 🚀 Production Deployment

### Binary Deployment
```bash
# Build binary
make build

# Set environment variables
export DB_HOST=your_production_db_host
export JWT_SECRET=your_production_jwt_secret
# ... other env vars

# Run binary
./bin/main
```

### Docker Deployment
```bash
# Build image
docker build -t blog-api .

# Run with environment file
docker run -p 8000:8000 --env-file .env blog-api
```

## 📄 License

MIT License

---

**Developer**: Syaiful  
**Repository**: [blog-app-api](https://github.com/Syaiful313/blog-app-api)