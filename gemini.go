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
	logger.Println("Creating a new AI client")
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

func aiChat(ctx context.Context, chatSession *genai.ChatSession, prompt string) (resp string, err error) {
	logger.Printf("Sending a message to Gemini %s", prompt)
	// Send the message to the generative model.
	r, err := chatSession.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	fnCalls := r.Candidates[0].FunctionCalls()

	// handle a conversation that had no calls, just
	if len(fnCalls) == 0 {
		logger.Printf("Non functional response received")
		for _, part := range r.Candidates[0].Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				resp += string(txt)
			}
		}
	} else {
		args := fnCalls[0].Args

		// the lookup of fucntions is invoked with the actual Dx function
		// and then the arguments are applied here.
		response := dxFns[fnCalls[0].Name](args)
		genAiResp := genai.FunctionResponse{
			Name:     fnCalls[0].Name,
			Response: response,
		}

		r, err = chatSession.SendMessage(ctx, genAiResp)
		if err != nil {
			return "", err
		}

		for _, part := range r.Candidates[0].Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				resp += string(txt)
			}
		}
	}

	return resp, err
}

var dxFns = map[string]DxFunc{
	"createGuest": createGuestDx,
	"allGuests":   allGuestsDx,
	"oneGuest":    oneGuestDx,
	"deleteGuest": deleteGuestDx,
	"updateGuest": updateGuestDx,
}

func aiTools() []*genai.Tool {
	return []*genai.Tool{
		{
			FunctionDeclarations: []*genai.FunctionDeclaration{
				{
					Name:        "createGuest",
					Description: "Create or register a new guest.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"name":  {Type: genai.TypeString},
							"email": {Type: genai.TypeString},
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
							"id": {Type: genai.TypeInteger},
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
							"id": {Type: genai.TypeInteger},
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
							"id":    {Type: genai.TypeInteger},
							"name":  {Type: genai.TypeString},
							"email": {Type: genai.TypeString},
						},
						Required: []string{"id", "name", "email"},
					},
				},
			},
		},
	}
}
