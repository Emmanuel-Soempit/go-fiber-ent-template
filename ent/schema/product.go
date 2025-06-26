package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.Float("price").
			Positive(),
		field.String("image_url").
			Optional(),
		field.String("category").
			NotEmpty().
			SchemaType(map[string]string{
				"mysql": "ENUM('plain', 'pocketed', 'drawstring', 'denim', 'velvet', 'aso_oke')",
			}),
		field.String("design").
			NotEmpty().
			SchemaType(map[string]string{
				"mysql": "ENUM('naomi', 'eden', 'snug', 'luxe_voyager', 'jubilee', 'salem', 'beulah', 'havilah', 'bethel', 'myrrh', 'tote_ayanfe')",
			}),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
