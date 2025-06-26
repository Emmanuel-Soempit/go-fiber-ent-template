package config

import (
	"context"
	"log"
	"os"
	"xaia-backend/ent"
	"xaia-backend/internal/api"
	"xaia-backend/internal/whatsapp/delivery/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitializeConfigurations() {
	// Load app .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to OS env")
	}
	log.Println(".env loaded successfully")

	app := fiber.New()
	appConfigurations(app)

	client := databseConfigs()
	// util.ToggleConversationalAutomation(true)
	// log.Println("Route /gd/webhookh pre registered")
	// app.Post("/gd/webhookh", func(ctx *fiber.Ctx) error {
	// 	log.Println("Whatsapp routes compil")
	// 	return ctx.SendStatus(fiber.StatusUnauthorized)
	// })
	// log.Println("Route /gd/webhookh registered")

	api.SetupRoutes(app, client)
	http.RegisterWhatsappRoutes(app, client)

	app.Listen(":3333")

}

func databseConfigs() *ent.Client {
	// Initialize database connection
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed opening connection to mysql: %v", err)
	}
	// defer client.Close()
	log.Println("Database Connected")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Migration Successful")

	return client
}

func appConfigurations(app *fiber.App) {
	// CORS should be first
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://whatsapp-vendor-frontend.vercel.app",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	log.Println("CORS middleware applied")

	// Then logger
	app.Use(logger.New())
	// Then your custom middleware
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request Origin: %s", c.Get("Origin"))
		return c.Next()
	})
	log.Println("App configurations successful")
}
