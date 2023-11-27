package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const FILENAME string = "gtp_instruction_config.json"

func main() {
	//preparedMessages := LoadConfiguration(FILENAME)
	var userInput UserInput

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "-l",
				Usage:       "programming language",
				Destination: &userInput.language,
			},
			&cli.BoolFlag{
				Name:        "-m",
				Usage:       "mutable/immutable",
				Destination: &userInput.mutable,
			},
			&cli.StringFlag{
				Name:        "-c",
				Usage:       "construct",
				Destination: &userInput.construct,
			},
			&cli.StringFlag{
				Name:        "-t",
				Usage:       "data type",
				Destination: &userInput.dataType,
			},
			&cli.StringFlag{
				Name:        "-u",
				Usage:       "unit",
				Destination: &userInput.unit,
			},
			&cli.StringFlag{
				Name:        "-d",
				Usage:       "description",
				Destination: &userInput.description,
			},
			&cli.IntFlag{
				Name:        "-n",
				Usage:       "number of suggestions",
				Destination: &userInput.numOfSuggestions,
			},
		},
		Name:  "naming",
		Usage: "give address pointers a senseful name",
		Action: func(*cli.Context) error {
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
