package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWriteTemplate(t *testing.T) {
	// Create a temporary test template file
	tempDir := t.TempDir()
	testFilePath := filepath.Join(tempDir, "test.html")

	// Create the test template content
	templateContent := `<html><body><h1>Hello {{.Role}}!</h1><p>{{.Content}}</p></body></html>`
	if err := os.WriteFile(testFilePath, []byte(templateContent), 0644); err != nil {
		t.Fatalf("Failed to create test template file: %v", err)
	}

	// Test data
	testData := &Message{
		Role:   User,
		Content: "This is a test!",
		Time:   time.Now(),
	}

	// Expected output
	expectedOutput := `<html><body><h1>Hello user!</h1><p>This is a test!</p></body></html>`

	// Create a buffer to capture output
	var buf bytes.Buffer

	// Execute the template
	err := writeChatMessage(&buf, testFilePath, testData)
	if err != nil {
		t.Fatalf("writeTemplate returned error: %v", err)
	}

	// Check the output
	if got := buf.String(); got != expectedOutput {
		t.Errorf("writeTemplate output = %q, want %q", got, expectedOutput)
	}
}

func TestBadTemplatePath(t *testing.T) {
	// Test with a bad file path
	badFilePath := "nonexistent.html"

	// Create a buffer to capture output
	var buf bytes.Buffer

	// Attempt to execute the template with a bad file path
	err := writeChatMessage(&buf, badFilePath, nil)
	if err == nil {
		t.Errorf("Expected error for bad file path, got nil")
	}
}
