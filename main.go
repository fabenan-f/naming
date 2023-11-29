package main

import (
	"log"
	"os"

	"github.com/dotenv-org/godotenvvault"
	"github.com/fabenan-f/naming/naming"
	openai "github.com/sashabaranov/go-openai"
)

const (
	FILENAME = "gtp_instruction_config.json"
)

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

func main() {
	apiKey := getApiKey()
	client := openai.NewClient(apiKey)

	preMessages := naming.ParseConfiguration(FILENAME)

	gptClient := naming.GtpClient{
		PreMessages: preMessages,
		Client:      *client,
	}

	app := naming.NewCliApp(gptClient)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
