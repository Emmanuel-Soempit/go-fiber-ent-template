package repository

import (
	"context"
	"xaia-backend/ent"
	"xaia-backend/ent/product"
	"xaia-backend/internal/api/product/delivery/http/dtos"
)

type entProductRepository struct {
	client *ent.Client
}

func NewEntProductRepository(client *ent.Client) ProductRepository {
	return &entProductRepository{client: client}
}

func (r *entProductRepository) Create(ctx context.Context, newProduct dtos.CreateProductRequest) (*ent.Product, error) {
	creator := r.client.Product.
		Create().
		SetName(newProduct.Name).
		SetPrice(newProduct.Price).
		SetCategory(newProduct.Category).
		SetDesign(newProduct.Design).
		SetDescription(newProduct.Description).
		SetImageURL(newProduct.ImageURL)

	if newProduct.Description != "" {
		creator.SetDescription(newProduct.Description)
	}

	if newProduct.ImageURL != "" {
		creator.SetImageURL(newProduct.ImageURL)
	}

	product, err := creator.Save(ctx)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *entProductRepository) Delete(ctx context.Context, id int) error {
	_, err := r.client.Product.Delete().Where(product.ID(id)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *entProductRepository) FindById(ctx context.Context, id int) (*ent.Product, error) {
	product, err := r.client.Product.Query().Where(product.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *entProductRepository) Update(ctx context.Context, id int, update dtos.UpdateProductRequest) (*ent.Product, error) {
	updater := r.client.Product.UpdateOneID(id)
	if update.Name != nil {
		updater.SetName(*update.Name)
	}
	if update.Description != nil {
		updater.SetDescription(*update.Description)
	}
	if update.Price != nil {
		updater.SetPrice(*update.Price)
	}
	if update.ImageURL != nil {
		updater.SetImageURL(*update.ImageURL)
	}
	if update.Category != nil {
		updater.SetCategory(*update.Category)
	}
	if update.Design != nil {
		updater.SetDesign(*update.Design)
	}
	product, err := updater.Save(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *entProductRepository) FindByCategoryAndDesign(ctx context.Context, category, design *string) ([]*ent.Product, error) {
	query := r.client.Product.Query()
	if category != nil {
		query = query.Where(product.CategoryEQ(*category))
	}
	if design != nil {
		query = query.Where(product.DesignEQ(*design))
	}
	return query.All(ctx)
}
