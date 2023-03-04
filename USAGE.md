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

// edits
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
	apiKey := "<your-api-key>"

    model := "text-davinci-edit-001"
    input := "What day of the wek is it?"
    instruction := "Fix the spelling mistakes"

    result, err := oacg.GetEdit(apiKey, model, input, instruction)
    if err != nil {
        fmt.Printf("Failed to get edit: %v\n", err)
        return
    }

    fmt.Printf("Edit result: %s\n", result)
}
```

// embeddings
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
    apiKey := "<your-api-key>"

    model := "text-embedding-ada-002"
    input := "The food was delicious and the waiter was very friendly."

    embeddings, err := oacg.GetEmbeddings(apiKey, model, input)
    if err != nil {
        fmt.Printf("Failed to get embeddings: %v\n", err)
        return
    }

    fmt.Printf("Embeddings: %v\n", embeddings)
}
```

// moderations
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
    apiKey := "<your-api-key>"

    input := "I want to kill them."

    score, err := oacg.GetModeration(apiKey, input)
    if err != nil {
        fmt.Printf("Failed to get moderation score: %v\n", err)
        return
    }

    fmt.Printf("Moderation score: %.2f\n", score)
}
```

// files
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
    apiKey := "<your-api-key>"

    // List all files
    fileList, err := oacg.ListFiles(apiKey)
    if err != nil {
        fmt.Printf("Failed to list files: %v\n", err)
        return
    }

    fmt.Println("List of files:")
    for _, file := range fileList {
        fmt.Printf("- %s (ID: %s)\n", file.Purpose, file.ID)
    }
}
```

// finetunes
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
    apiKey := "<your-api-key>"

	// Create a new fine-tune
	trainingFileID := "file-XGinujblHPwGLSztz8cPS8XY"
	fineTuneID, err := oacg.CreateFineTune(apiKey, trainingFileID)
	if err != nil {
		fmt.Printf("Failed to create fine-tune: %v\n", err)
		return
	}
	fmt.Printf("Fine-tune created with ID %s\n", fineTuneID)

	// List all fine-tunes
	fineTuneList, err := oacg.ListFineTunes(apiKey)
	if err != nil {
		fmt.Printf("Failed to list fine-tunes: %v\n", err)
		return
	}
	for _, fineTune := range fineTuneList {
		fmt.Printf("ID: %s, Model ID: %s, Status: %s\n", fineTune.ID, fineTune.ModelID, fineTune.Status)
	}

	// Get a specific fine-tune
	fineTune, err := oacg.GetFineTune(apiKey, fineTuneID)
	if err != nil {
		fmt.Printf("Failed to get fine-tune: %v\n", err)
		return
	}
	fmt.Printf("Fine-tune status: %s\n", fineTune.Status)

	// Cancel a fine-tune
	err = oacg.CancelFineTune(apiKey, fineTuneID)
	if err != nil {
		fmt.Printf("Failed to cancel fine-tune: %v\n", err)
		return
	}
	fmt.Println("Fine-tune cancelled")

	// Get fine-tune events
	events, err := oacg.GetFineTuneEvents(apiKey, fineTuneID)
	if err != nil {
		fmt.Printf("Failed to get fine-tune events: %v\n", err)
		return
	}
	for _, event := range events.Data {
		fmt.Printf("Event: %s, Status: %s, Percentage: %d, Created At: %d\n", event.Event, event.Status, event.Percentage, event.CreatedAt)
	}

	// Delete a fine-tune
	err = oacg.DeleteFineTune(apiKey, fineTuneID)
	if err != nil {
		fmt.Printf("Failed to delete fine-tune: %v\n", err)
		return
	}
	fmt.Println("Fine-tune deleted")
}
```

// transcribe audio
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
    apiKey := "<your-api-key>"

	audioFilePath := "./content/audio.mp3"
	model := "whisper-1"

	transcription, err := oacg.TranscribeAudio(apiKey, audioFilePath, model)
	if err != nil {
		panic(err)
	}

	fmt.Println(transcription.Text)
}
```

// translate audio
```go
package main

import (
	"fmt"

	"github.com/0xnu/oacg"
)

func main() {
    apiKey := "<your-api-key>"

	audioFilePath := "./content/german.m4a"
	model := "whisper-1"

	translation, err := oacg.TranslateAudio(apiKey, audioFilePath, model)
	if err != nil {
		panic(err)
	}

	fmt.Println(translation.Text)
}
```