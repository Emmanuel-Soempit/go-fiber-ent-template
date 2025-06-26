package usecase

import (
	"context"
	"xaia-backend/ent"
	"xaia-backend/internal/api/product/delivery/http/dtos"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req dtos.CreateProductRequest) (*ent.Product, error)
	DeleteProduct(ctx context.Context, productId int) error
	UpdateProduct(ctx context.Context, id int, req dtos.UpdateProductRequest) (*ent.Product, error)
	FindByCategoryAndDesign(ctx context.Context, category, design *string) ([]*ent.Product, error)
}
