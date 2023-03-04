For example, you can retrieve all the models and a model called `curie`.

// models
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
	apiKey := "<your-api-key>"

	// List all available models
	models, err := oacg.ListModels(apiKey)
	if err != nil {
		fmt.Printf("Failed to list models: %v\n", err)
		return
	}

	fmt.Println("Available Models:")
	for _, model := range models {
		fmt.Printf("Model Name: %s (%s): %s\n", model.ID, model.ModelType, model.Description)
	}

	// Get the "curie" model by ID
	modelID := "curie"
	model, err := oacg.GetModel(apiKey, modelID)
	if err != nil {
		fmt.Printf("Failed to get model with ID %s: %v\n", modelID, err)
		return
	}

	fmt.Printf("\nModel ID: %s, Model Type: %s, Description: %s\n", model.ID, model.ModelType, model.Description)
}
```


// completions
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
	apiKey := "<your-api-key>"

	model := "text-davinci-003"
	prompt := "This is a test"
	maxTokens := 7
	temperature := float32(0)

	result, err := oacg.GetCompletion(apiKey, model, prompt, maxTokens, temperature)
	if err != nil {
		fmt.Printf("Failed to get completion: %v\n", err)
		return
	}

	fmt.Printf("Completion result: %s\n", result)
}
```

// chat
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
	apiKey := "<your-api-key>"

	model := "gpt-3.5-turbo"
	messages := []oacg.ChatMessage{
		{Role: "user", Content: "Hello!"},
	}

	result, err := oacg.GetChatCompletion(apiKey, model, messages)
	if err != nil {
		fmt.Printf("Failed to get chat completion: %v\n", err)
		return
	}

	fmt.Printf("Chat completion result: %s\n", result)
}
```