package di

import (
	"fiber-starter/internal/handler"
	"fiber-starter/internal/repository"
	"fiber-starter/internal/service"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Container struct {
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
	PostHandler *handler.PostHandler
}

func NewContainer(db *pgxpool.Pool, jwtSecret string) *Container {
	// Repositories
	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db)
	postRepo := repository.NewPostRepository(db) 

	// Services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo, userRepo, jwtSecret)
	postService := service.NewPostService(postRepo) 

	// Handlers
	userHandler := handler.NewUserHandler(userService, authService)
	authHandler := handler.NewAuthHandler(authService)
	postHandler := handler.NewPostHandler(postService, authService) 

	return &Container{
		UserHandler: userHandler,
		AuthHandler: authHandler,
		PostHandler: postHandler, 
	}
}
