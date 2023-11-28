package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dotenv-org/godotenvvault"
	openai "github.com/sashabaranov/go-openai"
)

var apiKey string = getApiKey()
var client = openai.NewClient(apiKey)

func getApiKey() string {
	err := godotenvvault.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	api_key := os.Getenv("OPENAI_API_KEY")
	if api_key == "" {
		log.Fatal("OPENAI_API_KEY is not exported")
	}
	return api_key
}

func askGtpForName(messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
