package dtos

import (
	"time"
)

type WhatsAppWebhookPayload struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Value struct {
				MessagingProduct string `json:"messaging_product"`
				Metadata         struct {
					DisplayPhoneNumber string `json:"display_phone_number"`
					PhoneNumberID      string `json:"phone_number_id"`
				} `json:"metadata"`
				Contacts []struct {
					Profile struct {
						Name string `json:"name"`
					} `json:"profile"`
					WaID string `json:"wa_id"`
				} `json:"contacts"`
				Messages []struct {
					From      string `json:"from"`
					ID        string `json:"id"`
					Timestamp string `json:"timestamp"`
					Text      struct {
						Body string `json:"body"`
					} `json:"text"`
					Type        string `json:"type"`
					Interactive *struct {
						Type      string `json:"type"`
						ListReply *struct {
							ID    string `json:"id"`
							Title string `json:"title"`
						} `json:"list_reply,omitempty"`
						// Add button_reply if you use buttons
					} `json:"interactive,omitempty"`
				} `json:"messages"`
			} `json:"value"`
			Field string `json:"field"`
		} `json:"changes"`
	} `json:"entry"`
}

type Customer struct {
	WhatsAppID string    `json:"whatsapp_id"`    // e.g. the user’s E.164 number or Meta wa_id
	Name       *string   `json:"name,omitempty"` // Optional display name
	Phone      string    `json:"phone"`          // E.164 format
	CreatedAt  time.Time `json:"created_at"`     // When we first saw this user
	UpdatedAt  time.Time `json:"updated_at"`     // Last time we updated this record
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CustomerMetrics struct {
	PhoneNumberID      string    `json:"phone_number_id"`
	MessageID          string    `json:"message_id"`                     // e.g. WhatsApp’s message UUID
	From               string    `json:"from"`                           // user’s WhatsApp ID / E.164 number
	To                 string    `json:"to"`                             // your business number
	Timestamp          time.Time `json:"timestamp"`                      // when it was sent
	MessageType        string    `json:"message_type"`                   // "text", "image", "location", etc.
	TextBody           string    `json:"text_body,omitempty"`            // only for text messages
	BodySize           int       `json:"body_size,omitempty"`            // len(TextBody)
	MediaURLs          []string  `json:"media_urls,omitempty"`           // URLs to any media attachments
	InteractiveReplyID string    `json:"interactive_reply_id,omitempty"` // the id of the button/list row selected
	Location           *Location `json:"location,omitempty"`             // lat/long if a location message
	Name               string
}

type PromptListItems struct {
	Type   string           `json:"type"`
	Header PromptText       `json:"header"`
	Body   PromptText       `json:"body"`
	Footer PromptText       `json:"footer"`
	Action PromptListAction `json:"action"`
}

type PromptText struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text"`
}

type PromptListAction struct {
	Button   string              `json:"button"`
	Sections []PromptListSection `json:"sections"`
}

type PromptListSection struct {
	Title string       `json:"title"`
	Rows  []PromptItem `json:"rows"`
}

type PromptItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PromptNode struct {
	Type     string                `json:"type,omitempty"`
	Prompts  []string              `json:"prompts,omitempty"`
	Message  string                `json:"message,omitempty"`
	Items    *PromptListItems      `json:"items,omitempty"`
	Handler  string                `json:"handler,omitempty"` // <-- Add this line
	Children map[string]PromptNode `json:"children,omitempty"`
	Queries  []string              `json:"queries,omitempty"`
}
