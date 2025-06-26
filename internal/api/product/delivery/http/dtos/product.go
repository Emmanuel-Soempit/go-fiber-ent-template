package dtos

import "time"

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url,omitempty"`
	Category    string  `json:"category"` // allowed: plain, pocketed, drawstring, denim, velvet, aso_oke
	Design      string  `json:"design"`   // allowed: naomi, eden, snug, luxe_voyager, jubilee, salem, beulah, havilah, bethel, myrrh, tote_ayanfe
}

// Allowed values for Category and Design, matching the product schema enums.
var AllowedCategories = []string{"plain", "pocketed", "drawstring", "denim", "velvet", "aso_oke"}
var AllowedDesigns = []string{"naomi", "eden", "snug", "luxe_voyager", "jubilee", "salem", "beulah", "havilah", "bethel", "myrrh", "tote_ayanfe"}

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	Category    string    `json:"category"`
	Design      string    `json:"design"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	ImageURL    *string  `json:"image_url,omitempty"`
	Category    *string  `json:"category,omitempty"` // allowed: plain, pocketed, drawstring, denim, velvet, aso_oke
	Design      *string  `json:"design,omitempty"`   // allowed: naomi, eden, snug, luxe_voyager, jubilee, salem, beulah, havilah, bethel, myrrh, tote_ayanfe
}
