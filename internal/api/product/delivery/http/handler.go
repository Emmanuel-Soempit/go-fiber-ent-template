package http

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
	"xaia-backend/internal/api/product/delivery/http/dtos"
	"xaia-backend/internal/api/product/usecase"
	"xaia-backend/internal/util"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var req dtos.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate category
	validCategory := slices.Contains(dtos.AllowedCategories, req.Category)
	if !validCategory {
		return util.Failed(c, "Invalid category. Allowed: plain, pocketed, drawstring, denim, velvet, aso_oke", nil)
	}

	// Validate design
	validDesign := slices.Contains(dtos.AllowedDesigns, req.Design)
	if !validDesign {
		return util.Failed(c, "Invalid design. Allowed: naomi, eden, snug, luxe_voyager, jubilee, salem, beulah, havilah, bethel, myrrh, tote_ayanfe", nil)
	}
	log.Println("Parsed successfully:", req.Design)

	// Save file to public/images/
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		cleanFilename := strings.ReplaceAll(file.Filename, " ", "_")
		savePath := fmt.Sprintf("public/images/%d_%s", time.Now().UnixNano(), cleanFilename)
		if err := c.SaveFile(file, savePath); err != nil {
			return util.Failed(c, "Failed to save image", err.Error())
		}
		// Generate public URL (adjust if you serve static files differently)
		req.ImageURL = "/" + savePath
	}

	product, err := h.productUsecase.CreateProduct(c.Context(), req)
	if err != nil {
		log.Println("Create product error", err)
		return util.Failed(c, "Failed to create product", err.Error())
	}

	rsp := dtos.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageURL:    product.ImageURL,
		Category:    product.Category,
		Design:      product.Design,
		CreatedAt:   product.CreateTime,
		UpdatedAt:   product.UpdateTime,
	}

	return util.Created(c, "Product created successfully", rsp)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return util.Failed(c, "Invalid product ID", nil)
	}

	var req dtos.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return util.Failed(c, "Invalid request body", err.Error())
	}

	// Validate category and design if provided
	if req.Category != nil && !slices.Contains(dtos.AllowedCategories, *req.Category) {
		return util.Failed(c, "Invalid category. Allowed: plain, pocketed, drawstring, denim, velvet, aso_oke", nil)
	}
	if req.Design != nil && !slices.Contains(dtos.AllowedDesigns, *req.Design) {
		return util.Failed(c, "Invalid design. Allowed: naomi, eden, snug, luxe_voyager, jubilee, salem, beulah, havilah, bethel, myrrh, tote_ayanfe", nil)
	}

	file, err := c.FormFile("image")
	if err == nil && file != nil {
		cleanFilename := strings.ReplaceAll(file.Filename, " ", "_")
		savePath := fmt.Sprintf("public/images/%d_%s", time.Now().UnixNano(), cleanFilename)
		if err := c.SaveFile(file, savePath); err != nil {
			return util.Failed(c, "Failed to save image", err.Error())
		}
		// Generate public URL (adjust if you serve static files differently)
		imgURL := "/" + savePath
		req.ImageURL = &imgURL
	}

	product, err := h.productUsecase.UpdateProduct(c.Context(), id, req)
	if err != nil {
		return util.Failed(c, "Failed to update product", err.Error())
	}

	rsp := dtos.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageURL:    product.ImageURL,
		Category:    product.Category,
		Design:      product.Design,
		CreatedAt:   product.CreateTime,
		UpdatedAt:   product.UpdateTime,
	}

	return util.Success(c, "Product updated successfully", rsp)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return util.Failed(c, "Invalid product id", nil)
	}

	err = h.productUsecase.DeleteProduct(c.Context(), id)
	if err != nil {
		return util.Failed(c, "Faied to delete product", err.Error())
	}

	return util.Success(c, "Product deleted succesfully", nil)
}

func (h *ProductHandler) GetByCategoryAndDesign(c *fiber.Ctx) error {
	var categoryPtr, designPtr *string

	category := c.Query("category")
	if category != "" {
		if !slices.Contains(dtos.AllowedCategories, category) {
			return util.Failed(c, "Invalid category", nil)
		}
		categoryPtr = &category
	}

	design := c.Query("design")
	if design != "" {
		if !slices.Contains(dtos.AllowedDesigns, design) {
			return util.Failed(c, "Invalid design", nil)
		}
		designPtr = &design
	}

	products, err := h.productUsecase.FindByCategoryAndDesign(c.Context(), categoryPtr, designPtr)
	if err != nil {
		return util.Failed(c, "Failed to fetch products", err.Error())
	}

	var rsp []dtos.ProductResponse
	for _, product := range products {
		rsp = append(rsp, dtos.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			ImageURL:    product.ImageURL,
			Category:    product.Category,
			Design:      product.Design,
			CreatedAt:   product.CreateTime,
			UpdatedAt:   product.UpdateTime,
		})
	}

	return util.Success(c, "Products fetched successfully", rsp)
}
