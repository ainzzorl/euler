package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sort"
)

const NUM_SAMPLES = 1_000_000_000

// https://projecteuler.net/problem=587
// Concave triangle

//nolint:unusedfunc
func p587(ratio float64) int {
	samples := [NUM_SAMPLES]float64{}
	numSamples := 0
	for i := 0; numSamples < NUM_SAMPLES; i++ {
		x := rand.Float64() * 0.5
		y := rand.Float64() * 0.5
		distToCenter := math.Sqrt((x-0.5)*(x-0.5) + (y-0.5)*(y-0.5))
		if distToCenter < 0.5 {
			continue
		}
		samples[numSamples] = y / x
		numSamples++
	}

	sort.Float64s(samples[:numSamples])
	//	fmt.Printf("samples=%v\n", samples[:numSamples])
	index := int(float64(numSamples) * ratio)
	value := samples[index]
	estimate := int(math.Ceil(1 / value))
	fmt.Printf("index=%v, value=%v, estimate=%v\n", index, value, estimate)

	return estimate
}

// func main() {
// 	//fmt.Println(p587(0.1))
// 	fmt.Println(p587(0.001))
// }
