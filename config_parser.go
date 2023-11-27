package main

import (
	"encoding/json"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func LoadConfiguration(file string) []openai.ChatCompletionMessage {
	var prepareMessages []openai.ChatCompletionMessage
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&prepareMessages)
	return prepareMessages
}
