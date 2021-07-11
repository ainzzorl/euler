package main

import (
	"log"
	"math"
	"strconv"
)

// https://projecteuler.net/problem=080
//
// Square root digital expansion.
// Newton's algorithm.

// https://golang.org/pkg/math/big/ would do,
// but I felt like implementing my own.
// The implementation is very application-specific,
// not reusable.
type LongDecimal struct {
	digits []int
}

func (ld *LongDecimal) ToString() string {
	result := strconv.Itoa(ld.digits[0])
	result += "."
	for i := 1; i < len(ld.digits); i++ {
		result += strconv.Itoa(ld.digits[i])
	}
	return result
}

func (ld *LongDecimal) Add(other *LongDecimal) *LongDecimal {
	result := ld.clone()
	carry := 0

	for i := len(ld.digits) - 1; i >= 0; i-- {
		result.digits[i] += other.digits[i] + carry
		if i > 0 {
			carry = result.digits[i] / 10
			result.digits[i] %= 10
		}
	}

	return result
}

func (ld *LongDecimal) Substract(other *LongDecimal) *LongDecimal {
	result := ld.clone()
	carry := 0

	for i := len(ld.digits) - 1; i >= 0; i-- {
		result.digits[i] -= other.digits[i] + carry
		carry = 0
		for result.digits[i] < 0 {
			result.digits[i] += 10
			carry++
		}
	}

	if carry < 0 {
		log.Fatal("Unexpected carry at the end of long substraction")
	}

	return result
}

func (ld *LongDecimal) Divide(other *LongDecimal) *LongDecimal {
	result := Make(0, len(ld.digits))

	cur := ld.clone()

	for i := 0; i < len(ld.digits); i++ {
		for m := 1; ; m++ {
			mm := other.shortMult(m)
			if !mm.lessOrEqual(cur) {
				result.digits[i] = m - 1
				cur = cur.Substract(other.shortMult(m - 1))
				cur = cur.shortMult(10)
				break
			}
		}
	}

	return result
}

func Sqrt(s int, numDigits int, conversionPrecision int) *LongDecimal {
	current := Make(1, numDigits)

	for {
		prev := current.clone()
		current = current.shortDivide(2).Add(Make(s, numDigits).Divide(prev).shortDivide(2))

		eq := true
		for j := 0; j < conversionPrecision; j++ {
			if current.digits[j] != prev.digits[j] {
				eq = false
			}
		}
		if eq {
			return current
		}
	}
}

func (ld *LongDecimal) shortMult(other int) *LongDecimal {
	result := ld.clone()
	carry := 0

	for i := len(ld.digits) - 1; i >= 0; i-- {
		result.digits[i] = result.digits[i]*other + carry
		if i > 0 {
			carry = result.digits[i] / 10
			result.digits[i] %= 10
		}
	}

	return result
}

func (ld *LongDecimal) shortDivide(other int) *LongDecimal {
	result := Make(0, len(ld.digits))
	carry := 0

	for i := 0; i < len(ld.digits); i++ {
		result.digits[i] = ld.digits[i] + carry*10
		carry = result.digits[i] % other
		result.digits[i] /= other
	}

	return result
}

func (ld *LongDecimal) lessOrEqual(other *LongDecimal) bool {
	for i := 0; i < len(ld.digits); i++ {
		if ld.digits[i] < other.digits[i] {
			return true
		}
		if ld.digits[i] > other.digits[i] {
			return false
		}
	}
	return true
}

func (ld *LongDecimal) clone() *LongDecimal {
	digits := make([]int, len(ld.digits))
	copy(digits, ld.digits)
	return &LongDecimal{
		digits: digits,
	}
}

func p080() int {
	result := 0
	for i := 2; i < 100; i++ {
		if isPerfectSquare(i) {
			continue
		}
		sqrt := Sqrt(i, 140, 120)
		for j := 0; j < 100; j++ {
			result += sqrt.digits[j]
		}
	}
	return result
}

func Make(from int, numDigits int) *LongDecimal {
	digits := make([]int, numDigits)
	digits[0] = from
	return &LongDecimal{
		digits: digits,
	}
}

func isPerfectSquare(no int) bool {
	int_root := int(math.Sqrt(float64(no)))
	return int_root*int_root == no
}
