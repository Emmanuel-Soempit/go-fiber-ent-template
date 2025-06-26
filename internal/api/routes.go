package api

import (
	"xaia-backend/ent"
	authHttp "xaia-backend/internal/api/auth/delivery/http"
	productHttp "xaia-backend/internal/api/product/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, client *ent.Client) {
	app.Static("/public", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routeGroup := app.Group("/api/v1")
	authHttp.RegisterAuthRoutes(routeGroup, client)
	productHttp.RegisterProductRoutes(routeGroup, client)
}
