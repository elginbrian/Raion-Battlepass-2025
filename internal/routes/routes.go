package routes

import (
	"raion-battlepass/config"
	"raion-battlepass/internal/handler"
	"raion-battlepass/internal/middleware"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	postHandler *handler.PostHandler,
	jwtSecret string,
) {
	config.InitMetrics()

	app.Use(config.PrometheusMiddleware)

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	app.Get("/api", redirectToDocs)
	app.Get("/docs", redirectToDocs)
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	setupUserRoutes(app, userHandler)
	setupAuthRoutes(app, authHandler, jwtSecret)
	setupPostRoutes(app, postHandler)
	setupSearchRoutes(app, userHandler, postHandler)
}

func redirectToDocs(c *fiber.Ctx) error {
	return c.Redirect("/docs/index.html")
}

func setupSearchRoutes(app *fiber.App, userHandler *handler.UserHandler, postHandler *handler.PostHandler) {
	searchGroup := app.Group("/api/search")
	searchGroup.Get("/users", userHandler.SearchUsers)
	searchGroup.Get("/posts", postHandler.SearchPosts)
}

func setupUserRoutes(app *fiber.App, handler *handler.UserHandler) {
	userGroup := app.Group("/api/users")
	userGroup.Put("/", handler.UpdateUser)
	userGroup.Get("/", handler.GetAllUsers)
	userGroup.Get("/:id", handler.GetUserByID)
}

func setupAuthRoutes(app *fiber.App, handler *handler.AuthHandler, jwtSecret string) {
	authGroup := app.Group("/api/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
	authGroup.Get("/current-user", handler.GetUserInfo, middleware.TokenValidationMiddleware(jwtSecret))
	authGroup.Put("/change-password", handler.ChangePassword, middleware.TokenValidationMiddleware(jwtSecret))
}

func setupPostRoutes(app *fiber.App, handler *handler.PostHandler) {
	postGroup := app.Group("/api/posts")
	postGroup.Post("/", handler.CreatePost)
	postGroup.Put("/:id", handler.UpdatePost)
	postGroup.Delete("/:id", handler.DeletePost)
	postGroup.Get("/", handler.GetAllPosts)
	postGroup.Get("/:id", handler.GetPostByID)
	postGroup.Get("/user/:user_id", handler.GetPostsByUserID)
}