package main

import (
	"testing"
)

func TestP549(t *testing.T) {
	actual := p549(100)
	if actual != 2012 {
		t.Errorf("P549 failed for n=%v\n", 2012)
		t.Errorf("Expected: %v\n", 2012)
		t.Errorf("Actual: %v\n", actual)
	}
}
