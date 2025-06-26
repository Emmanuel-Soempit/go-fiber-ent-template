// ent/schema/customer.go

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the WhatsApp user.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		// The WhatsApp platform user ID (e.g. the phone in E.164, or Meta’s wa_id)
		field.String("whatsapp_id").
			Unique().
			NotEmpty().
			Comment("Unique identifier for the user on WhatsApp"),

		// Optional human‑readable name (if available via webhook profile)
		field.String("name").
			Optional().
			Comment("User’s display name from WhatsApp profile"),

		// E.164 phone number (same as whatsapp_id in many cases, but stored separately if needed)
		field.String("phone").
			NotEmpty().
			Comment("User’s phone number in E.164 format"),

		// // A JSON blob for any credentials or tokens (e.g. access token, refresh token)
		// field.JSON("credentials", map[string]interface{}{}).
		//     Optional().
		//     Comment("Platform‑specific credentials or session data"),

		// Timestamps
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Customer (if you later relate messages, sessions, etc.)
func (Customer) Edges() []ent.Edge {
	return nil
}
