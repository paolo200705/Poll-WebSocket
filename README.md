# 🗳️ Live WebSocket Polling System

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.10.1-green.svg)](https://gin-gonic.com/)
[![WebSocket](https://img.shields.io/badge/WebSocket-Real--time-orange.svg)](https://github.com/gorilla/websocket)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A high-performance, real-time polling system built with Go, featuring WebSocket connections for instant vote updates and dynamic poll management. Perfect for live events, surveys, and interactive presentations.

## ✨ Key Features

### 🚀 **Real-Time Voting**
- **Instant Updates**: WebSocket-powered live vote counting with zero refresh needed
- **Multi-User Support**: Concurrent users can vote simultaneously with real-time synchronization
- **Vote Tracking**: Individual user vote mapping with JWT-based authentication

### 🔧 **Dynamic Poll Management**
- **Live Option Editing**: Add or remove poll options while voting is active
- **Flexible Poll Creation**: Create polls with custom questions and multiple options
- **UUID-Based Identification**: Secure, unique poll identification system

### 🔐 **Secure Authentication**
- **JWT Token System**: Stateless authentication with configurable expiration
- **Auto-Generated User IDs**: Seamless user identification without registration
- **Protected Endpoints**: All poll operations require valid authentication

### 🌐 **RESTful API + WebSocket**
- **Hybrid Architecture**: REST API for poll management + WebSocket for real-time features
- **Clean JSON Responses**: Well-structured API responses for easy integration
- **CORS Support**: Cross-origin requests enabled for web applications

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Client App    │◄──►│   Gin Router     │◄──►│  Poll Storage   │
│  (Web/Mobile)   │    │  + Middleware    │    │  (In-Memory)    │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         │              ┌──────────────────┐             │
         └──────────────►│  WebSocket Hub   │◄────────────┘
                        │ (Real-time Sync) │
                        └──────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- **Go 1.24.5+** installed on your system
- **Git** for cloning the repository

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Otavio-Fina/live-websocket.git
   cd live-websocket
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   # Copy the example .env file
   cp .env.example .env
   
   # Edit .env with your JWT secret
   JWT_SECRET=your-super-secret-jwt-key-here
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080` 🎉

## 📖 API Usage

### Authentication

First, get your authentication token:

```bash
curl -X GET http://localhost:8080/auth/login
```

**Response:**
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Poll Management

#### Create a Poll
```bash
curl -X POST http://localhost:8080/poll \
  -H "token: YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Favorite Programming Language",
    "question": "Which programming language do you prefer?",
    "options": ["Go", "Python", "JavaScript", "Rust"]
  }'
```

#### Get All Polls
```bash
curl -X GET http://localhost:8080/poll \
  -H "token: YOUR_JWT_TOKEN"
```

#### Get Specific Poll
```bash
curl -X GET "http://localhost:8080/poll/POLL_ID" \
  -H "token: YOUR_JWT_TOKEN"
```

### WebSocket Real-Time Features

Connect to a poll's WebSocket endpoint:
```
ws://localhost:8080/ws/poll/POLL_ID?token=YOUR_JWT_TOKEN
```

#### WebSocket Message Types

**1. Initialize Connection**
```json
{
  "mensagge_type": 1
}
```

**2. Cast Vote**
```json
{
  "mensagge_type": 2,
  "vote": "Go"
}
```

**3. Modify Poll Options (Live)**
```json
{
  "mensagge_type": 3,
  "change_options_params": {
    "TypeScript": "add",
    "COBOL": "del"
  }
}
```

## 🛠️ Development

### Project Structure

```
live-websocket/
├── controller/          # Business logic and WebSocket handlers
│   └── controller.go
├── middleware/          # Authentication and request processing
│   └── middleware.go
├── models/             # Data structures and global state
│   └── model.go
├── routes/             # HTTP route handlers
│   └── routes.go
├── bruno/              # API testing collection (Bruno)
│   └── poll/
├── main.go             # Application entry point
├── go.mod              # Go module dependencies
└── .env                # Environment configuration
```

### Key Dependencies

- **[Gin](https://gin-gonic.com/)**: High-performance HTTP web framework
- **[Gorilla WebSocket](https://github.com/gorilla/websocket)**: WebSocket implementation
- **[JWT-Go](https://github.com/golang-jwt/jwt)**: JSON Web Token authentication
- **[UUID](https://github.com/google/uuid)**: Unique identifier generation
- **[GoDotEnv](https://github.com/joho/godotenv)**: Environment variable loading

### Running Tests

API tests are available in the `bruno/` directory. Install [Bruno](https://www.usebruno.com/) to run the test collection:

```bash
# Install Bruno CLI
npm install -g @usebruno/cli

# Run API tests
bru run bruno/poll
```

### Development Commands

```bash
# Run with hot reload (install air first)
go install github.com/cosmtrek/air@latest
air

# Format code
go fmt ./...

# Run linter
golangci-lint run

# Build for production
go build -o bin/live-websocket main.go
```

## 🤝 Contributing

We welcome contributions! Here's how you can help:

### Development Workflow

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**
4. **Add tests** for new functionality
5. **Ensure code quality**
   ```bash
   go fmt ./...
   go vet ./...
   golangci-lint run
   ```
6. **Commit with clear messages**
   ```bash
   git commit -m "feat: add real-time poll analytics"
   ```
7. **Push and create a Pull Request**

### Coding Standards

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions focused and testable
- Handle errors appropriately

### Areas for Contribution

- 📊 **Analytics Dashboard**: Real-time vote analytics and charts
- 🗄️ **Database Integration**: Replace in-memory storage with persistent DB
- 🔒 **Enhanced Security**: Rate limiting, input validation, HTTPS
- 📱 **Mobile SDK**: Native mobile app integration
- 🎨 **Admin Interface**: Web-based poll management UI
- 🧪 **Testing**: Unit tests and integration tests
- 📚 **Documentation**: API documentation and tutorials

## 🔮 Roadmap

- [ ] **Database Persistence** - PostgreSQL/MongoDB integration
- [ ] **Poll Analytics** - Vote statistics and real-time charts  
- [ ] **User Management** - Registration, profiles, and permissions
- [ ] **Poll Templates** - Pre-built poll types and themes
- [ ] **Export Features** - CSV/PDF result exports
- [ ] **Webhook Support** - External system integrations
- [ ] **Mobile Apps** - iOS and Android native applications

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Gin Framework** team for the excellent HTTP framework
- **Gorilla WebSocket** contributors for robust WebSocket support
- **Go Community** for the amazing ecosystem and tools

## 📞 Support

- 🐛 **Bug Reports**: [Create an issue](https://github.com/Otavio-Fina/live-websocket/issues)
- 💡 **Feature Requests**: [Start a discussion](https://github.com/Otavio-Fina/live-websocket/discussions)
- 📧 **Contact**: [Your Email](mailto:your-email@example.com)

---

<div align="center">

**Built with ❤️ using Go and WebSockets**

[⭐ Star this repo](https://github.com/Otavio-Fina/live-websocket) if you find it useful!

</div>