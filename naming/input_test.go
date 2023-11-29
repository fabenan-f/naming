package naming

import (
	"testing"
)

func TestTransformToInstruction(t *testing.T) {
	input := Input{
		language:         "Java",
		symbol:           "variable",
		description:      "it should represent a cat that is at the same time a dog",
		numOfSuggestions: 10,
	}

	want := "I have a Java application. Give me a name for a variable. The variable should represent following: it should represent a cat that is at the same time a dog. Give me 10 suggestions"
	got := input.transformToInstruction()
	if got != want {
		t.Fatalf(`TransformToInstruction()="%s" is not "%s"`, got, want)
	}
}

func TestValidateInput(t *testing.T) {
	input := Input{
		numOfSuggestions: 10,
	}

	err := input.validateInput()
	if err != nil {
		t.Fatalf(`%v number of suggestions should not cause an error since =< 10`, input.numOfSuggestions)
	}
}

func TestValidateInputError(t *testing.T) {
	input := Input{
		numOfSuggestions: 11,
	}

	err := input.validateInput()
	if err == nil {
		t.Fatalf(`%v number of suggestions should cause an error since > 10`, input.numOfSuggestions)
	}
}
