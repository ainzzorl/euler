package main

import (
	"testing"
)

func Test068(t *testing.T) {
	result := p068()
	expected := "6531031914842725"

	if result != expected {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n", result)
	}
}
