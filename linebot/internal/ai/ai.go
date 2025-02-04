package ai

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type AiService struct {
	client *genai.Client
	model  string
}

func NewAiService(client *genai.Client, model string) *AiService {
	return &AiService{
		client: client,
		model:  model,
	}
}

func (s *AiService) CallGemini(ctx context.Context, message string) (string, error) {
	systemPrompt, err := os.ReadFile("prompt.txt")
	if err != nil {
		return "", fmt.Errorf("failed to read prompt file: %w", err)
	}

	// Generate content
	resp, err := s.client.Models.GenerateContent(ctx, s.model, genai.Text(message), &genai.GenerateContentConfig{
		SystemInstruction: genai.Text(string(systemPrompt))[0],
	})
	if err != nil {
		return "", fmt.Errorf("generate content err: %w", err)
	}

	content := resp.Candidates[0].Content.Parts[0].Text

	return string(content), nil
}

func NewCreateExpenseFunctionCall() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "create_expense",
		Description: "Create an expense of user.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"category": {
					Type:        genai.TypeString,
					Description: "Category of the expense. Like 'Food', 'Entertainment', 'Clothing', 'Transportation', etc.",
				},
				"description": {
					Type:        genai.TypeString,
					Description: "Detail description. Let user know what the expense is for. Not too long.",
				},
				"currency": {
					Type:        genai.TypeString,
					Description: "Currency of the expense. If not specified, default to TWD.",
				},
				"amount": {
					Type:        genai.TypeNumber,
					Description: "Amount of the expense.",
				},
			},
			Required: []string{"category", "description", "currency", "amount"},
		},
	}
}
