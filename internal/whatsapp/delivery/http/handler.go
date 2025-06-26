package http

import (
	"fmt"
	"log"
	"os"
	"xaia-backend/internal/whatsapp/client"
	"xaia-backend/internal/whatsapp/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	handler usecase.BotUsecase
}

func VerifyWebhook(c *fiber.Ctx) error {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	expectedToken := os.Getenv("WHATSAPP_VERIFICATION_TOKEN")

	if mode == "subscribe" && token == expectedToken {
		fmt.Println("Check complete")
		return c.SendString(challenge)
	}

	fmt.Println("Check Failed")
	return c.SendStatus(fiber.StatusForbidden)
}

func (h *Handler) WebhookHandler(ctx *fiber.Ctx) error {

	log.Println("Whatsapp bot triggered")
	clientInstance := client.NewClient(ctx)
	h.handler.HandlePrompt(ctx.Context(), clientInstance)
	return ctx.SendStatus(fiber.StatusNoContent)
}
