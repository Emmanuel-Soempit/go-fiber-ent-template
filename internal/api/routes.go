package api

import (
	"xaia-backend/ent"
	"xaia-backend/internal/api/auth/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, client *ent.Client) {
	routeGroup := app.Group("/api/v1")

	http.RegisterAuthRoutes(routeGroup, client)
}
