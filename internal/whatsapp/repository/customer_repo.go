package repository

import (
	"context"
	"xaia-backend/ent"
	"xaia-backend/internal/whatsapp/delivery/http/dtos"
)

type CustomerRepo interface {
	FindByWhatsappId(ctx context.Context, whatsappId string) (*ent.Customer, error)
	CreateNewCustomer(ctx context.Context, customer dtos.Customer) (*ent.Customer, error)
}
