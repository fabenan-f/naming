package naming

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type GtpClient struct {
	PreMessages []openai.ChatCompletionMessage
	Client      openai.Client
}

func (c *GtpClient) askForSuggestions(message openai.ChatCompletionMessage) (string, error) {
	messages := append(c.PreMessages, message)

	resp, err := c.Client.CreateChatCompletion(
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
