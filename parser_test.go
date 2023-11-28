package main

import (
	"fmt"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	input := "gtp_instruction_config.json"

	want := "Imagine you are a software engineer looking for a fitting name"
	got := ParseConfiguration(input)
	if want != got[0].Content {
		t.Fatalf(`LoadConfiguration(%q)=%q is not %q`, input, got[0].Content, want)
	}
}

func TestParseResponse(t *testing.T) {
	inputResponse := "turtleSpeedInSeconds,turtleSpeedPerSecond"
	inputNumOfSuggestions := 2

	want := "\x1b[37m1. \x1b[32mturtleSpeedInSeconds\n\x1b[37m2. \x1b[32mturtleSpeedPerSecond\n"
	got, _ := ParseResponse(inputResponse, inputNumOfSuggestions)

	if got != want {
		t.Fatalf(`ParseResponse(%q, %v)=%q is not %q`, inputResponse, inputNumOfSuggestions, got, want)
	}

	fmt.Println(got)
}

func TestParseResponseError(t *testing.T) {
	inputResponse := "I need more details to give an answer."
	inputNumOfSuggestions := 2

	_, err := ParseResponse(inputResponse, inputNumOfSuggestions)

	if err == nil {
		t.Fatalf(`Expected an error message for gtp response: %s`, inputResponse)
	}
}
