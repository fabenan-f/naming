package main

import (
	"fmt"
	"strings"
)

type input struct {
	language         string
	mutable          bool
	symbol           string
	dataType         string
	unit             string
	description      string
	numOfSuggestions int
}

func (input *input) transformToInstruction() string {
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
