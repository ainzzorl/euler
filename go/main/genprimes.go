package main

// GeneratePrimes generates prime numbers.
func GeneratePrimes(max int) []int {
	visited := make([]bool, max+1)
	result := []int{}
	for i := 2; i <= max; i++ {
		if visited[i] {
			continue
		}
		result = append(result, i)
		for j := i * 2; j <= max; j += i {
			visited[j] = true
		}
	}
	return result
}
