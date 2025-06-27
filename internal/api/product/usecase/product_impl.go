package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"xaia-backend/ent"
	"xaia-backend/internal/api/product/delivery/http/dtos"
	"xaia-backend/internal/api/product/repository"
)

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) CreateProduct(c context.Context, req dtos.CreateProductRequest) (*ent.Product, error) {
	return u.productRepo.Create(c, req)
}

func (u *productUsecase) DeleteProduct(ctx context.Context, productId int) error {
	product, err := u.productRepo.FindById(ctx, productId)
	if err != nil {
		return err
	}

	imagePath := strings.TrimPrefix(product.ImageURL, "/")

	err = os.Remove(imagePath)
	if err != nil {
		log.Printf("%s image url:%s", err, imagePath)
		return errors.New("error deleting product image")
	}

	return u.productRepo.Delete(ctx, productId)
}

func (u *productUsecase) UpdateProduct(ctx context.Context, id int, req dtos.UpdateProductRequest) (*ent.Product, error) {

	//Remove previous image
	if req.ImageURL != nil {
		product, err := u.productRepo.FindById(ctx, id)
		if err != nil {
			return nil, err
		}

		var imagePath string
		if os.Getenv("Environment") == "local" {
			imagePath = strings.TrimPrefix(product.ImageURL, "/")
		} else {
			imagePath = fmt.Sprintf("..%s", product.ImageURL)
		}

		err = os.Remove(imagePath)
		if err != nil {
			log.Printf("%s image url:%s", err, imagePath)
			return nil, errors.New("error deleting product image")
		}
	}

	return u.productRepo.Update(ctx, id, req)
}

func (u *productUsecase) FindByCategoryAndDesign(ctx context.Context, category, design *string) ([]*ent.Product, error) {
	return u.productRepo.FindByCategoryAndDesign(ctx, category, design)
}
