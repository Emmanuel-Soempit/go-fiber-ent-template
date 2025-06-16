package config

import (
	"context"
	"log"
	"os"
	"xaia-backend/ent"
	"xaia-backend/internal/api"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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

	api.SetupRoutes(app, client)

	app.Listen(":3000")

}

func databseConfigs() *ent.Client {
	// Initialize database connection
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL"))
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
	// Initialize default config
	app.Use(logger.New())
}
