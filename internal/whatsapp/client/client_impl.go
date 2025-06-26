package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"xaia-backend/internal/whatsapp/delivery/http/dtos"

	"github.com/gofiber/fiber/v2"
)

type client struct {
	ctx      *fiber.Ctx
	customer dtos.CustomerMetrics
}

func NewClient(ctx *fiber.Ctx) Client {
	return &client{ctx: ctx}
}
func (c *client) ParseClientPayload() (dtos.CustomerMetrics, error) {
	var customerMetrics dtos.CustomerMetrics
	var payload dtos.WhatsAppWebhookPayload

	if err := c.ctx.BodyParser(&payload); err != nil {
		return customerMetrics, err
	}

	if len(payload.Entry) == 0 ||
		len(payload.Entry[0].Changes) == 0 ||
		len(payload.Entry[0].Changes[0].Value.Messages) == 0 {
		return customerMetrics, nil
	}

	entry := payload.Entry[0]
	change := entry.Changes[0]
	value := change.Value
	message := value.Messages[0]

	customerMetrics.PhoneNumberID = value.Metadata.PhoneNumberID
	customerMetrics.MessageID = message.ID
	customerMetrics.From = message.From
	customerMetrics.To = value.Metadata.DisplayPhoneNumber
	if ts, err := parseTimestamp(message.Timestamp); err == nil {
		customerMetrics.Timestamp = ts
	}
	customerMetrics.MessageType = message.Type

	// Add this block to extract the customer's name
	if len(value.Contacts) > 0 {
		customerMetrics.Name = value.Contacts[0].Profile.Name
	}

	if message.Type == "text" {
		log.Println("Text")
		customerMetrics.TextBody = message.Text.Body
		customerMetrics.BodySize = len(message.Text.Body)
	} else if message.Type == "interactive" && message.Interactive != nil {
		if message.Interactive.Type == "list_reply" && message.Interactive.ListReply != nil {
			log.Println("Interactive: list_reply")
			log.Printf("Text: %s", message.Interactive.ListReply.ID)
			customerMetrics.TextBody = message.Interactive.ListReply.ID
			customerMetrics.InteractiveReplyID = message.Interactive.ListReply.ID
		}
	}
	// TODO: request_welcome webhokk handler

	c.customer = customerMetrics
	return customerMetrics, nil
}

// Helper to parse WhatsApp timestamp string to time.Time
func parseTimestamp(ts string) (time.Time, error) {
	// WhatsApp usually sends Unix timestamp as string
	sec, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
}

func (c *client) SendMessage(resType string, value interface{}) error {
	switch resType {
	case "text":
		text, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string for text message, got %T", value)
		}
		return c.sendTextMessage(text)
	case "list":
		var listData map[string]interface{}
		switch v := value.(type) {
		case map[string]interface{}:
			listData = v
		default:
			// Try to convert struct to map
			m, err := structToMap(v)
			if err != nil {
				return fmt.Errorf("could not convert list struct to map: %w", err)
			}
			listData = m
		}
		return c.sendListMessage(listData)
	default:
		return fmt.Errorf("unsupported message type: %s", resType)
	}
}

func (c *client) sendTextMessage(message string) error {
	phoneID := os.Getenv("WHATSAPP_PHONE_ID")
	token := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", phoneID)
	log.Printf("Message to send: %s to %s", message, c.customer.From)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                c.customer.From,
		"type":              "text",
		"text": map[string]string{
			"body": message,
		},
	}

	return c.doRequest(url, token, payload)
}

func (c *client) sendListMessage(listData map[string]interface{}) error {
	phoneID := os.Getenv("WHATSAPP_PHONE_ID")
	token := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", phoneID)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                c.customer.From,
		"type":              "interactive",
		"interactive":       listData,
	}
	log.Printf("Send list message: %s", listData)

	return c.doRequest(url, token, payload)
}

func (c *client) doRequest(url, token string, payload map[string]interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("error from http: %s", string(respBytes))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to send message: %s", resp.Status)
	}
	return nil
}

func structToMap(v interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *client) SendProduct(imageURL, caption string) error {
	phoneID := os.Getenv("WHATSAPP_PHONE_ID")
	token := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", phoneID)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                c.customer.From,
		"type":              "image",
		"image": map[string]string{
			"link":    imageURL,
			"caption": caption, // WhatsApp supports captions for images
		},
	}

	return c.doRequest(url, token, payload)
}
