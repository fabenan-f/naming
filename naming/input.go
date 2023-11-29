package naming

import (
	"errors"
	"fmt"
	"strings"
)

type Input struct {
	language         string
	mutable          bool
	symbol           string
	dataType         string
	unit             string
	description      string
	numOfSuggestions int
}

func (input *Input) transformToInstruction() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "I have a %s application. Give me a name for a ", input.language)
	if input.mutable {
		builder.WriteString("mutable ")
	}
	builder.WriteString(input.symbol)

	if input.dataType != "" || input.unit != "" {
		fmt.Fprint(&builder, " which")
	}

	if input.dataType != "" {
		fmt.Fprintf(&builder, " data has type %s", input.dataType)
	}

	if input.dataType != "" && input.unit != "" {
		builder.WriteString(" and")
	}

	if input.unit != "" {
		fmt.Fprintf(&builder, " the unit %s", input.unit)
	}

	fmt.Fprintf(&builder, ". The %s should represent following: %s. Give me %v suggestions",
		input.symbol, input.description, input.numOfSuggestions)

	return builder.String()
}

func (input *Input) validateInput() error {
	if input.numOfSuggestions > 10 {
		return errors.New(
			fmt.Sprintf(
				`Maximal number of suggestion is 10, you wanted %v`,
				input.numOfSuggestions,
			),
		)
	}
	return nil
}
