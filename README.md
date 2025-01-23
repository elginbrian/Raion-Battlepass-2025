# 🦁 Raion Battlepass 2025 (Backend)

This is a **RESTful API** for a simple social media application. It allows users to manage their posts, including creating, updating, and deleting posts, and provides authentication using JWT. The API is built using the Fiber framework and interacts with a PostgreSQL database.

## 🚀 Features

- **User Authentication:** Secure user authentication using JWT.
- **Post Management:** Create, update, delete, and retrieve posts with ease.
- **Scalable Architecture:** Designed for scalability and maintainability.
- **Monitoring:** Integrated with Prometheus for metrics and monitoring.
- **API Documentation:** Swagger UI for exploring and testing API endpoints.

## 🛠️ Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white) ![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white) ![DigitalOcean](https://img.shields.io/badge/DigitalOcean-0080FF?style=for-the-badge&logo=digitalocean&logoColor=white) ![Prometheus](https://img.shields.io/badge/Prometheus-0098FF?style=for-the-badge&logo=prometheus&logoColor=white) ![Swagger](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=white) ![Vitepress](https://img.shields.io/badge/Vitepress-3EAF8D?style=for-the-badge&logo=vite&logoColor=white)

- **Backend Framework:** [Go](https://golang.org/) with [Fiber](https://gofiber.io/)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Containerization:** [Docker](https://www.docker.com/)
- **Hosting:** [DigitalOcean](https://www.digitalocean.com/)
- **Monitoring:** [Prometheus](https://prometheus.io/)
- **Documentation:** [Swagger](https://swagger.io/) and [Vitepress](https://vitepress.vuejs.org/)

## 📂 Project Structure

```plaintext
.
├── cmd
│   ├── index.html
│   └── main.go
├── config
│   ├── config.go
│   ├── metrics.go
│   └── migration.go
├── db
│   └── migrations/
├── internal
│   ├── di/
│   ├── domain/
│   ├── handler/
│   ├── middleware/
│   ├── repository/
│   ├── routes/
│   └── service/
├── pkg
│   ├── request/
│   └── response/
├── public
│   └── uploads/
├── vendor/
├── vitepress
│   ├── Dockerfile
│   ├── docs
│   ├── index.md
│   ├── package-lock.json
│   ├── package.json
│   ├── sticker.png
│   └── vite.config.js
├── Dockerfile
├── docker-compose.yaml
├── docs/
├── go.mod
├── go.sum
├── prometheus.yaml
├── README.md
└── wait-for-it.sh
```

## 🏁 Getting Started

### Prerequisites

- [Go](https://golang.org/) installed
- [Docker](https://www.docker.com/) installed
- [PostgreSQL](https://www.postgresql.org/) database

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/elginbrian/Raion-Battlepass-2025.git
   cd Raion-Battlepass-2025
   ```
2. Set up the environment variables:

   > Copy the .env.example file to .env and configure the required variables.

3. Start the application:
   ```bash
   docker compose build
   docker compose up
   ```
4. Access the application via:
   ```bash
   localhost:8084
   ```
