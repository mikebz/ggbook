package main

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// NOTE: we are using a singleton for this simple example
// but for a production system you need to do proper model management
var model *genai.GenerativeModel

func createAiClient(ctx context.Context, apiKey string) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func configureAiModel(client *genai.Client) {
	model = client.GenerativeModel("gemini-1.5-flash")
	model.Tools = aiTools()
}

func aiTools() []*genai.Tool {
	t := genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{
			{
				Name:        "createGuest",
				Description: "Create or register a new guest.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"name": {
							Type: genai.TypeString,
						},
						"email": {
							Type: genai.TypeString,
						},
					},
					Required: []string{"name", "email"},
				},
			},
			{
				Name:        "allGuests",
				Description: "Get all the guests that have been created or registered.",
			},
			{
				Name:        "oneGuest",
				Description: "Get one guest",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"id": {
							Type: genai.TypeInteger,
						},
					},
					Required: []string{"id"},
				},
			},
			{
				Name:        "deleteGuest",
				Description: "Delete a guest.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"id": {
							Type: genai.TypeInteger,
						},
					},
					Required: []string{"id"},
				},
			},
			{
				Name:        "updateGuest",
				Description: "Update guest information such as name and email",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"id": {
							Type: genai.TypeInteger,
						},
						"name": {
							Type: genai.TypeString,
						},
						"email": {
							Type: genai.TypeString,
						},
					},
					Required: []string{"id", "name", "email"},
				},
			},
		},
	}

	return []*genai.Tool{&t}
}
