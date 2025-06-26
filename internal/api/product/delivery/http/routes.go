package http

import (
	"xaia-backend/ent"
	"xaia-backend/internal/api/product/repository"
	"xaia-backend/internal/api/product/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(router fiber.Router, client *ent.Client) {
	productRepo := repository.NewEntProductRepository(client)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := NewProductHandler(productUsecase)

	productRoutes := router.Group("/products")
	{
		productRoutes.Post("", productHandler.Create)
		productRoutes.Delete("/:id", productHandler.Delete)
		productRoutes.Put("/:id", productHandler.Update)
		productRoutes.Get("", productHandler.GetByCategoryAndDesign)
	}
}
