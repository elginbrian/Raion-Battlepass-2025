# Fiber Starter

Fiber Starter is a clean architecture starter project for building scalable and maintainable web applications using [Fiber](https://gofiber.io/), a Go web framework inspired by Express.js.

## Directory Structure

```
fiber-starter/
├── .github/
│   └── workflows/
│       └── ci-cd.yaml
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── db/
│   └── migrations/
│       ├── 000001_init_users.up.sql
│       ├── 000001_init_users.down.sql
├── internal/
│   ├── di/
│   │   └── di.go
│   ├── domain/
│   │   └── entity.go
│   ├── handler/
│   │   ├── user_handler.go
│   │   └── auth_handler.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   └── auth_repository.go
│   ├── routes/
│   │   └── routes.go
│   ├── service/
│   │   ├── user_service.go
│   │   └── auth_service.go
├── pkg/
│   └── response/
│       └── response.go
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
└── README.md
```

## How to Run the Project

### Running Locally

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/fiber-starter.git
   cd fiber-starter
   ```

2. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set Up Environment Variables**:
   Create a `.env` file in the root directory with the appropriate environment variables.

4. **Run Database Migrations**:

   ```bash
   migrate -path ./db/migrations -database "$DATABASE_URL" up
   ```

5. **Run the Application**:

   ```bash
   go run ./cmd/main.go
   ```

6. **Access the Application**:
   Open your browser and navigate to `http://localhost:3000`.

### Running with Docker

1. **Build and Run the Containers**:

   ```bash
   docker-compose up --build
   ```

2. **Access the Application**:
   Open your browser and navigate to `http://localhost:3000`.

### Running Tests

Run the tests using the following command:

```bash
go test ./...
```
