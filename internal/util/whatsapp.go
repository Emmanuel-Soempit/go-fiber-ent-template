package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"xaia-backend/internal/whatsapp/delivery/http/dtos"
)

// toggleConversationalAutomation enables or disables the welcome message
// for your WhatsApp Business phone number.
//
//	condition = true  → enable
//	condition = false → disable
func ToggleConversationalAutomation(condition bool) error {
	phoneNumberID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	if phoneNumberID == "" {
		return fmt.Errorf("WHATSAPP_PHONE_NUMBER_ID is not set")
	}
	accessToken := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	if accessToken == "" {
		return fmt.Errorf("WHATSAPP_ACCESS_TOKEN is not set")
	}

	// Build the request URL
	url := fmt.Sprintf(
		"https://graph.facebook.com/v17.0/%s/conversational_automation?enable_welcome_message=%t",
		phoneNumberID,
		condition,
	)

	// Create the POST request
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	// Add auth header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Execute
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("calling API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from API", resp.StatusCode)
	}

	return nil
}

func LoadPrompts(filename string) (dtos.PromptNode, error) {
	data, err := os.ReadFile(filename) // Reads the file into a byte slice
	if err != nil {
		return dtos.PromptNode{}, err
	}

	var prompts dtos.PromptNode
	if err := json.Unmarshal(data, &prompts); err != nil {
		return dtos.PromptNode{}, err
	}

	return prompts, nil
}

func MakeProductListURL(category, design string) string {
	return fmt.Sprintf("%s/products?category=%s&design=%s", os.Getenv("FRONTEND_BASE_URL"), category, design)
}

func MakePublicURL(imageURL string) string {

	return fmt.Sprintf("https://0a2d-105-112-117-5.ngrok-free.app%s", imageURL)
}
