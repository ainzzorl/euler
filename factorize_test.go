package euler

import (
	"reflect"
	"testing"
)

func TestFactorize(t *testing.T) {
	testFactorize(t, 2, []FactorizationEntry{{2, 1}})
	testFactorize(t, 10, []FactorizationEntry{{2, 1}, {5, 1}})
	testFactorize(t, 32, []FactorizationEntry{{2, 5}})
	testFactorize(t, 239, []FactorizationEntry{{239, 1}})
}

func testFactorize(t *testing.T, n int, expected []FactorizationEntry) {
	primes := GeneratePrimes(n)
	actual := Factorize(n, primes)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Factorize failed for n=%v\n", n)
		t.Errorf("Expected: %v\n", expected)
		t.Errorf("Actual: %v\n", actual)
	}
}
