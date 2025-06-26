package repository

import (
	"context"
	"xaia-backend/ent"
	"xaia-backend/internal/whatsapp/delivery/http/dtos"
)

type entCustomerRepo struct {
	client *ent.Client
}

func NewCustomerRepo(client *ent.Client) CustomerRepo {
	return &entCustomerRepo{client: client}
}

func (c *entCustomerRepo) FindByWhatsappId(ctx context.Context, whatsappId string) (*ent.Customer, error) {
	return nil, nil
}

func (c *entCustomerRepo) CreateNewCustomer(ctx context.Context, customer dtos.Customer) (*ent.Customer, error) {
	return nil, nil
}
