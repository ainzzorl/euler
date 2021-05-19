package main

import (
	"testing"
)

func TestP107(t *testing.T) {
	actual := p107("../data/p107_test.txt")
	if actual != 150 {
		t.Errorf("P107 failed\n")
		t.Errorf("Expected: %v\n", 150)
		t.Errorf("Actual: %v\n", actual)
	}
}
