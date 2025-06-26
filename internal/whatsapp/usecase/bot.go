package usecase

import (
	"context"
	"xaia-backend/internal/whatsapp/client"
)

type BotUsecase interface {
	HandlePrompt(ctx context.Context, client client.Client) error
}
