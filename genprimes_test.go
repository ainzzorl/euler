package main

import (
	"reflect"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	testGeneratePrimes(t, 1, []int{})
	testGeneratePrimes(t, 2, []int{2})
	testGeneratePrimes(t, 10, []int{2, 3, 5, 7})
	testGeneratePrimes(t, 11, []int{2, 3, 5, 7, 11})
}

func testGeneratePrimes(t *testing.T, max int, expected []int) {
	actual := GeneratePrimes(max)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("GeneratePrimes failed for max=%v\n", max)
		t.Errorf("Expected: %v\n", expected)
		t.Errorf("Actual: %v\n", actual)
	}
}
