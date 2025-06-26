package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strings"
	"xaia-backend/internal/api/product/usecase"
	"xaia-backend/internal/util"
	"xaia-backend/internal/whatsapp/client"
	"xaia-backend/internal/whatsapp/delivery/http/dtos"
	"xaia-backend/internal/whatsapp/repository"
)

type botUsecase struct {
	customerRepo   repository.CustomerRepo
	productUsecase usecase.ProductUsecase
	prompts        dtos.PromptNode
	handlers       map[string]func(context.Context, client.Client, *dtos.PromptNode, dtos.CustomerMetrics) error
}

func NewBotUsecase(customerRepo repository.CustomerRepo, productUsecase usecase.ProductUsecase) BotUsecase {
	prompts, err := util.LoadPrompts("data/prompts.json")
	if err != nil {
		log.Printf("Error %s", err)
		return nil
	}
	log.Println("Prompts loaded successfully")
	handlers := map[string]func(context.Context, client.Client, *dtos.PromptNode, dtos.CustomerMetrics) error{
		"welcomeHandler": welcomeHandler,
		// Add more handlers here
	}
	return &botUsecase{
		customerRepo:   customerRepo,
		productUsecase: productUsecase,
		prompts:        prompts,
		handlers:       handlers,
	}
}

func ResponseToString(resp interface{}) (string, error) {
	switch v := resp.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	default:
		// try to marshal it to JSON
		b, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("cannot convert response to string: %w", err)
		}
		return string(b), nil
	}
}

func (b *botUsecase) HandlePrompt(ctx context.Context, client client.Client) error {
	in, err := client.ParseClientPayload()
	if err != nil {
		return err
	}

	node, query := findPromptNode(b.prompts, in.TextBody)

	if node == nil {
		// Optionally send a default "I didn't understand" message
		return nil
	}
	log.Printf("Found prompt node: %s", node.Handler)

	if query != "" {
		return b.productHandler(ctx, client, query)
	}
	log.Printf("Found Query: %s", query)

	// If a handler is specified, call it
	if node.Handler != "" {
		log.Println("Handling node with handler")
		if handlerFunc, ok := b.handlers[node.Handler]; ok {
			return handlerFunc(ctx, client, node, in)
		}
	}

	log.Println("Fallback triggered")
	// Fallback to default behavior
	switch node.Type {
	case "text":
		client.SendMessage("text", node.Message)
	case "list":
		client.SendMessage("list", node.Items)
	default:
		// Handle other types if needed
	}
	return nil
}

// secondary functions
func findPromptNode(node dtos.PromptNode, input string) (*dtos.PromptNode, string) {
	if slices.Contains(node.Prompts, input) {
		return &node, ""
	}
	for _, query := range node.Queries {
		if input == query {
			return &node, query
		}
	}
	for _, child := range node.Children {
		if found, q := findPromptNode(child, input); found != nil {
			return found, q
		}
	}
	return nil, ""
}

func welcomeHandler(ctx context.Context, client client.Client, node *dtos.PromptNode, in dtos.CustomerMetrics) error {
	return client.SendMessage("text", fmt.Sprintf(node.Message, in.Name))
}

func (b *botUsecase) productHandler(ctx context.Context, client client.Client, query string) error {
	parts := strings.SplitN(query, "_", 2)
	if len(parts) != 2 {
		return client.SendMessage("text", "Sorry, I couldn't find that product.")
	}
	category := parts[0]
	design := parts[1]

	products, err := b.productUsecase.FindByCategoryAndDesign(ctx, &category, &design)
	if err != nil || len(products) == 0 {
		return client.SendMessage("text", "Sorry, I couldn't find any products for that selection.")
	}

	maxToSend := 2
	for i, product := range products {
		if i >= maxToSend {
			break
		}
		imageURL := util.MakePublicURL(product.ImageURL)
		caption := fmt.Sprintf("%s\n%s\nPrice: %v", product.Name, product.Description, product.Price)
		err := client.SendProduct(imageURL, caption)
		if err != nil {
			log.Printf("Failed to send product %s: %v", product.Name, err)
		}
	}

	if len(products) > maxToSend {
		link := util.MakeProductListURL(category, design)
		msg := fmt.Sprintf("There are more designs under this category! Click the link to view all: %s \n\nOr you can choose another category: \n/categories", link)
		client.SendMessage("text", msg)
	}
	return nil
}
