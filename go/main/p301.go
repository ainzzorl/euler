package main

// https://projecteuler.net/problem=301
// Nim

import (
	"fmt"
)

func p301(max int) int {
	result := 0
	for i := 1; i <= max; i++ {
		if i%1000000 == 0 {
			fmt.Printf("%v\n", i)
		}
		nimSum := i ^ (2 * i) ^ (3 * i)
		if nimSum == 0 {
			result++
		}
	}
	return result
}
