package http

import (
	"xaia-backend/ent"
	"xaia-backend/internal/api/auth/repository"
	"xaia-backend/internal/api/auth/usecase"
	"xaia-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(router fiber.Router, client *ent.Client) {

	authGroup := router.Group("/auth")

	authGroup.Use(middleware.CheckJwtToken)

	userRepo := repository.NewEntUserRepo(client)

	authUsecase := usecase.NewAuthUsecase(userRepo)

	authHandler := &Handler{authUc: authUsecase}

	// Routes
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/register", authHandler.Register)

}
