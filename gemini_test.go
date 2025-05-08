package main

import (
	"context"
	"os"
	"testing"

	"github.com/google/generative-ai-go/genai"
	"github.com/stretchr/testify/assert"
)


func TestCreateAiClient(t *testing.T) {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := createAiClient(ctx, apiKey)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	client.Close()
}

func TestConfigureAiModel(t *testing.T) {
	ctx := context.Background()

	// setting up the AI related objects, the step by step checks are done in a special test
	client, err := createAiClient(ctx, os.Getenv("GEMINI_API_KEY"))
	assert.NoError(t, err)
	defer client.Close()
	configureAiModel(client)
	assert.NotNil(t, model)
}

func TestAiTools(t *testing.T) {
	tools := aiTools()
	assert.Len(t, tools, 1)
	assert.GreaterOrEqual(t, len(tools[0].FunctionDeclarations), 5)
}

// This test ensures we get a simple function result
// back with no parameters.  It's not using our internal
// interpretation but rather checks wether the call is identified
// by the base layers.
func TestAllGuestPrompt(t *testing.T) {
	ctx := context.Background()

	// setting up the AI related objects, the step by step checks are done in a special test
	client, err := createAiClient(ctx, os.Getenv("GEMINI_API_KEY"))
	assert.NoError(t, err)
	defer client.Close()
	configureAiModel(client)
	session := model.StartChat()

	prompt := "I'd like to get a list of all users"
	// Send the message to the generative model.
	resp, err := session.SendMessage(ctx, genai.Text(prompt))
	assert.NoError(t, err)

	fnCalls := resp.Candidates[0].FunctionCalls()
	assert.GreaterOrEqual(t, len(fnCalls), 1)
	assert.Equal(t, fnCalls[0].Name, "allGuests")
}

// This test ensures we get a simple function result
// with a couple of parameters.  Not using `aiChat` but
// making sure that the base gemini service identifies the calls
// correctly.
func TestCreateGuestPrompt(t *testing.T) {
	ctx := context.Background()

	// setting up the AI related objects, the step by step checks are done in a special test
	client, err := createAiClient(ctx, os.Getenv("GEMINI_API_KEY"))
	assert.NoError(t, err)
	defer client.Close()
	configureAiModel(client)
	session := model.StartChat()

	prompt := "I'd like to register a guest named Mike with email hello@test.com"
	// Send the message to the generative model.
	resp, err := session.SendMessage(ctx, genai.Text(prompt))
	assert.NoError(t, err)

	fnCalls := resp.Candidates[0].FunctionCalls()
	assert.GreaterOrEqual(t, len(fnCalls), 1)
	assert.Equal(t, fnCalls[0].Name, "createGuest")
	assert.Len(t, fnCalls[0].Args, 2)
	assert.Equal(t, fnCalls[0].Args["name"], "Mike")
	assert.Equal(t, fnCalls[0].Args["email"], "hello@test.com")
}

func TestChat(t *testing.T) {
	ctx := context.Background()

	// setting up the AI related objects, the step by step checks are done in a special test
	client, err := createAiClient(ctx, os.Getenv("GEMINI_API_KEY"))
	assert.NoError(t, err)
	defer client.Close()
	configureAiModel(client)
	session := model.StartChat()

	response, err := aiChat(ctx, session, "give me a list of all guests")
	assert.NoError(t, err)
	assert.Contains(t, response, "list of all")
}

func TestNoCallMessage(t *testing.T) {
	ctx := context.Background()

	// setting up the AI related objects, the step by step checks are done in a special test
	client, err := createAiClient(ctx, os.Getenv("GEMINI_API_KEY"))
	assert.NoError(t, err)
	defer client.Close()
	configureAiModel(client)
	session := model.StartChat()

	response, err := aiChat(ctx, session, "what are you capable of doing?")
	assert.NoError(t, err)
	assert.Contains(t, response, "guest")
}
