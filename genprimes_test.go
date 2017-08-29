package euler

import (
	"reflect"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	doTest(t, 1, []int{})
	doTest(t, 2, []int{2})
	doTest(t, 10, []int{2, 3, 5, 7})
	doTest(t, 11, []int{2, 3, 5, 7, 11})
}

func doTest(t *testing.T, max int, expected []int) {
	actual := GeneratePrimes(max)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("GeneratePrimes failed for max=%v\n", max)
		t.Errorf("Expected: %v\n", expected)
		t.Errorf("Actual: %v\n", actual)
	}
}
