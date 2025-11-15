package main

// https://projecteuler.net/problem=090
// Cube digit pairs

import (
	"fmt"
)

func p090() int {
	result := 0
	total := 0
	// countForOne([]int{1, 2, 3, 4, 6, 7})
	for i1 := 0; i1 < 10; i1++ {
		for i2 := i1 + 1; i2 < 10; i2++ {
			for i3 := i2 + 1; i3 < 10; i3++ {
				for i4 := i3 + 1; i4 < 10; i4++ {
					for i5 := i4 + 1; i5 < 10; i5++ {
						for i6 := i5 + 1; i6 < 10; i6++ {
							for j1 := 0; j1 < 10; j1++ {
								for j2 := j1 + 1; j2 < 10; j2++ {
									for j3 := j2 + 1; j3 < 10; j3++ {
										for j4 := j3 + 1; j4 < 10; j4++ {
											for j5 := j4 + 1; j5 < 10; j5++ {
												for j6 := j5 + 1; j6 < 10; j6++ {
													total++
													if isValid([]int{i1, i2, i3, i4, i5, i6}, []int{j1, j2, j3, j4, j5, j6}) {
														result++
														//														fmt.Printf("(%v, %v) is valid\n", []int{i1, i2, i3, i4, i5, i6}, []int{j1, j2, j3, j4, j5, j6})
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	//fmt.Printf("total=%v, result=%v\n", total, result)
	return result / 2
}

func isValid(digits1 []int, digits2 []int) bool {
	has1 := []bool{false, false, false, false, false, false, false, false, false, false}
	for _, digit := range digits1 {
		if digit == 9 {
			has1[6] = true
		} else {
			has1[digit] = true
		}
	}
	has2 := []bool{false, false, false, false, false, false, false, false, false, false}
	for _, digit := range digits2 {
		if digit == 9 {
			has2[6] = true
		} else {
			has2[digit] = true
		}
	}

	squares := [][]int{
		{0, 1},
		{0, 4},
		{0, 6},
		{1, 6},
		{2, 5},
		{3, 6},
		{4, 6},
		{6, 4},
		{8, 1},
	}
	for _, square := range squares {
		h12 := has1[square[0]] && has2[square[1]]
		h21 := has2[square[0]] && has1[square[1]]
		if !h12 && !h21 {
			//fmt.Printf("(%v, %v) is not satisfiable by %v\n", digits1, digits2, square)
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(p090())
}
