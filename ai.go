package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/openai/openai-go/v1"
    "net/http"
)

func main() {
    // Set up the OpenAI API client with your API key
    apiKey := "your-api-key-here"
    client, err := openai.NewClient(apiKey)
    if err != nil {
        panic(err)
    }

    // Set up the Gin web server
    r := gin.Default()

    // Define a route to serve the index page
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    // Define a route to handle the text generation request
    r.POST("/generate", func(c *gin.Context) {
        // Get the prompt from the form data
        prompt := c.PostForm("prompt")

        // Set up the parameters for the text generation request
        model := "text-davinci-002"
        params := &openai.CompletionRequest{
            Prompt: prompt,
            Model:  model,
            Temperature: 0.7,
            MaxTokens: 60,
        }

        // Send the text generation request to Open
        
        resp, err := client.Completions.Create(params)
        if err != nil {
            c.String(http.StatusInternalServerError, err.Error())
            return
        }

        // Get the generated text from the response
        generatedText := resp.Choices[0].Text

        // Render the generated text on the page
        c.HTML(http.StatusOK, "index.html", gin.H{
            "prompt":        prompt,
            "generatedText": generatedText,
        })
    })

    // Serve static files from the "static" directory
    r.Static("/static", "./static")

   }
