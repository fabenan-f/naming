package main

import "testing"

func TestLoadConfiguration(t *testing.T) {
	input := "gtp_instruction_config.json"

	want := "Imagine you are software engineer looking for a fitting name"
	got := LoadConfiguration(input)
	if want != got[0].Content {
		t.Fatalf(`LoadConfiguration(%q)=%q is not %q`, input, got[0].Content, want)
	}
}
