package main

// FactorizationEntry Factorization entry.
type FactorizationEntry struct {
	p     int
	count int
}

// Factorize Factorize a number given primes.
func Factorize(n int, primes []int) []FactorizationEntry {
	result := []FactorizationEntry{}
	resultIndex := -1
	for primeIndex := 0; n > 1; {
		if n%primes[primeIndex] == 0 {
			n /= primes[primeIndex]
			if resultIndex >= 0 && result[resultIndex].p == primes[primeIndex] {
				result[resultIndex].count++
			} else {
				result = append(result, FactorizationEntry{primes[primeIndex], 1})
				resultIndex++
			}
		} else {
			if primes[primeIndex]*primes[primeIndex] > n {
				result = append(result, FactorizationEntry{n, 1})
				break
			} else {
				primeIndex++
			}
		}
	}
	return result
}
