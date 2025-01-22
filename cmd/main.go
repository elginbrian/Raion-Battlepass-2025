package main

import (
	"context"
	"log"
	"raion-battlepass/config"
	"raion-battlepass/internal/di"
	"raion-battlepass/internal/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/jackc/pgx/v4/pgxpool"

	_ "raion-battlepass/docs"
)

// @title RAION BATTLEPASS API
// @version 1.0
// @description This is a RESTful API for a simple social media application. It allows users to manage their posts, including creating, updating, and deleting posts, and provides authentication using JWT. The API is built using the Fiber framework and interacts with a PostgreSQL database.
// @termsOfService http://swagger.io/terms/

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host raion-battlepass.elginbrian.com
// @BasePath /
func main() {
	serverPort := config.GetServerPort()
	databaseURL := config.GetDatabaseURL()
	jwtSecret := config.GetJWTSecret()

	db, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Running database migrations...")
	if err := config.RunSQLMigrations(db); err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	container := di.NewContainer(db, jwtSecret)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRoutes(app, container.UserHandler, container.AuthHandler, container.PostHandler, jwtSecret)

	app.Static("/uploads", "./public/uploads")
	app.Static("/", "./cmd")

	log.Printf("Server is running on port %s", serverPort)
	if err := app.Listen(serverPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}