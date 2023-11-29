package naming

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	openai "github.com/sashabaranov/go-openai"
)

func ParseConfiguration(file string) []openai.ChatCompletionMessage {
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

func parseResponse(response string, numOfSuggestions int) (string, error) {
	var output strings.Builder

	suggestions := strings.Split(removeWhiteSpace(response), ",")
	if len(suggestions) != numOfSuggestions {
		return "", errors.New("Unexpected response: " + response)
	}

	output.WriteString("\n")
	for index, suggestion := range suggestions {
		output.WriteString(fmt.Sprintf("\033[37m%v. \033[32m%s\n", index+1, suggestion))
	}

	return output.String(), nil
}

func removeWhiteSpace(str string) string {
	var strWithoutSpace strings.Builder
	strWithoutSpace.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			strWithoutSpace.WriteRune(ch)
		}
	}
	return strWithoutSpace.String()
}
