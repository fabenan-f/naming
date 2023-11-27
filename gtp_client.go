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
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("OPENAI_API_KEY")
}

func askGtpForName(request string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: request,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
