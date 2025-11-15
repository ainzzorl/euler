package main

import (
	"testing"
)

func TestP090(t *testing.T) {
	actual := p090()
	if actual != 1217 {
		t.Errorf("P090 failed\n")
		t.Errorf("Expected: %v\n", 1217)
		t.Errorf("Actual: %v\n", actual)
	}
}
