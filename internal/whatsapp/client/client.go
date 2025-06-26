package client

import (
	"xaia-backend/internal/whatsapp/delivery/http/dtos"
)

type Client interface {
	ParseClientPayload() (dtos.CustomerMetrics, error)
	SendMessage(resType string, value interface{}) error
	SendProduct(imageURL, caption string) error
}
