package main

import (
    "fmt"
    "github.com/openai/openai-go/v1"
)

func main() {
    // Set up the OpenAI API client with your API key
    apiKey := "your-api-key-here"
    client, err := openai.NewClient(apiKey)
    if err != nil {
        panic(err)
    }

    // Set up the parameters for the text generation request
    prompt := "Once upon a time"
    model := "text-davinci-002"
    params := &openai.CompletionRequest{
        Prompt: prompt,
        Model:  model,
        Temperature: 0.7,
        MaxTokens: 60,
    }

    // Send the text generation request to OpenAI
    resp, err := client.Completions.Create(params)
    if err != nil {
        panic(err)
    }

    // Print the generated text
    fmt.Println(resp.Choices[0].Text)
}
