package main

// https://projecteuler.net/problem=321
// Swapping Counters
//
// See http://oeis.org/A006451

func p321(numTerms int) int {
	a := make([]int, numTerms+2)
	a[0] = 1
	a[1] = 1
	a[2] = 2
	a[3] = 4
	for i := 4; i <= numTerms+1; i++ {
		a[i] = 6*a[i-2] - a[i-4]
	}
	result := 0
	for i := 2; i <= numTerms+1; i++ {
		result += a[i] - 1
	}
	return result
}
