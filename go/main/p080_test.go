package main

import (
	"testing"
)

func Test080_Add(t *testing.T) {
	a := LongDecimal{
		digits: []int{1, 2, 3, 4, 5},
	}
	b := LongDecimal{
		digits: []int{2, 9, 8, 7, 6},
	}
	result := a.Add(&b).ToString()
	expected := "4.2221"

	if result != expected {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n", result)
	}
}

func Test080_Substract(t *testing.T) {
	a := LongDecimal{
		digits: []int{2, 9, 8, 7, 6},
	}
	b := LongDecimal{
		digits: []int{1, 2, 3, 4, 9},
	}
	result := a.Substract(&b).ToString()
	expected := "1.7527"

	if result != expected {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n", result)
	}
}

func Test080_Divide(t *testing.T) {
	a := LongDecimal{
		digits: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
	}
	b := LongDecimal{
		digits: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	result := a.Divide(&b).ToString()
	expected := "8.00000007"

	if result != expected {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n", result)
	}
}

func Test080_Sqrt(t *testing.T) {
	result := Sqrt(2, 5, 5).ToString()
	expected := "1.4142"

	if result != expected {
		t.Errorf("Expected: %s\n", expected)
		t.Errorf("Actual: %s\n", result)
	}
}

func Test080_SqrtSumDigits(t *testing.T) {
	sqrt2 := Sqrt(2, 140, 120)
	result := 0
	for i := 0; i < 100; i++ {
		result += sqrt2.digits[i]
	}
	expected := 475

	if result != expected {
		t.Errorf("Expected: %d\n", expected)
		t.Errorf("Actual: %d\n", result)
	}
}

func Test080_Actual(t *testing.T) {
	result := p080()
	expected := 40886

	if result != expected {
		t.Errorf("Expected: %d\n", expected)
		t.Errorf("Actual: %d\n", result)
	}
}
