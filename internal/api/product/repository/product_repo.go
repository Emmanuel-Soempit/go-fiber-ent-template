package repository

import (
	"context"
	"xaia-backend/ent"
	"xaia-backend/internal/api/product/delivery/http/dtos"
)

type ProductRepository interface {
	Create(ctx context.Context, newProduct dtos.CreateProductRequest) (*ent.Product, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (*ent.Product, error)
	Update(ctx context.Context, id int, update dtos.UpdateProductRequest) (*ent.Product, error)
	FindByCategoryAndDesign(ctx context.Context, category, design *string) ([]*ent.Product, error)
}
