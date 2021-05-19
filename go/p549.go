package main

// https://projecteuler.net/problem=549
// Divisibility of factorials

import (
	"fmt"
	"time"
)

func p549(n int) int {
	primes := GeneratePrimes(n)
	result := 0
	startTime := time.Now()
	intervalStartTime := time.Now()
	for i := 2; i <= n; i++ {
		minFactorial := 0
		for _, factorizationEntry := range Factorize(i, primes) {
			cur := p549ForEntry(factorizationEntry)
			if cur > minFactorial {
				minFactorial = cur
			}
		}
		result += minFactorial
		if i%1000 == 0 {
			intervalElapsed := time.Since(intervalStartTime)
			totalElapsed := time.Since(startTime)
			intervalStartTime = time.Now()
			fmt.Printf("n=%v, last 1000 processed in %v, total %v\n", i, intervalElapsed, totalElapsed)
		}
	}
	return result
}

func p549ForEntry(entry FactorizationEntry) int {
	total := 0
	for i := entry.p; ; i += entry.p {
		for j := i; j%entry.p == 0; j /= entry.p {
			total++
		}
		if total >= entry.count {
			return i
		}
	}
}
