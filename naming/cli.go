package naming

import (
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
)

func NewCliApp(client GtpClient) *cli.App {
	var input Input
	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Aliases:     []string{"l"},
				Usage:       "programming language",
				Destination: &input.language,
				Required:    true,
			},
			&cli.BoolFlag{
				Name:        "mutable",
				Aliases:     []string{"m"},
				Usage:       "mutable",
				Destination: &input.mutable,
			},
			&cli.StringFlag{
				Name:        "symbol",
				Aliases:     []string{"s"},
				Usage:       "identifier for variable, function, class etc. names",
				Destination: &input.symbol,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "type",
				Aliases:     []string{"t"},
				Usage:       "data type",
				Destination: &input.dataType,
			},
			&cli.StringFlag{
				Name:        "unit",
				Aliases:     []string{"u"},
				Usage:       "unit",
				Destination: &input.unit,
			},
			&cli.StringFlag{
				Name:        "description",
				Aliases:     []string{"d"},
				Usage:       "description",
				Destination: &input.description,
				Required:    true,
			},
			&cli.IntFlag{
				Name:        "num",
				Aliases:     []string{"n"},
				Usage:       "number of suggestions",
				Destination: &input.numOfSuggestions,
				Required:    true,
			},
		},
		Name:  "naming",
		Usage: "give symbols a meaningful name and make thereby your code more maintainable",
		Action: func(*cli.Context) error {
			err := input.validateInput()
			if err != nil {
				log.Fatal(err)
			}

			instruction := input.transformToInstruction()
			message := openai.ChatCompletionMessage{
				Role:    "user",
				Content: instruction,
			}
			response, err := client.askForSuggestions(message)
			if err != nil {
				log.Fatal(err)
			}

			parsedResponse, err := parseResponse(response, input.numOfSuggestions)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(parsedResponse)
			return nil
		},
	}
}
