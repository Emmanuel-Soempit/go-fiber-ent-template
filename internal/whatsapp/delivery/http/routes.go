package http

import (
	// "log"
	"xaia-backend/ent"

	"xaia-backend/internal/whatsapp/repository"
	"xaia-backend/internal/whatsapp/usecase"

	productRepo "xaia-backend/internal/api/product/repository"
	productUsecase "xaia-backend/internal/api/product/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterWhatsappRoutes(router *fiber.App, client *ent.Client) {

	customerRepo := repository.NewCustomerRepo(client)
	productRepo := productRepo.NewEntProductRepository(client)
	productUsecase := productUsecase.NewProductUsecase(productRepo)
	bot := usecase.NewBotUsecase(customerRepo, productUsecase)
	whatsappHandler := &Handler{handler: bot}

	router.Post("/webhook", whatsappHandler.WebhookHandler)
	// func(ctx *fiber.Ctx) error {
	// 	log.Println("Whatsapp routes compiled wrongly")
	// 	return ctx.SendStatus(fiber.StatusUnauthorized)
	// }
	router.Get("/webhook", VerifyWebhook)

}
