package main

import "fmt"

// https://projecteuler.net/problem=068

func p068() string {
	result := ""
	for i1 := 1; i1 < 10; i1++ {
		for i2 := 1; i2 < 10; i2++ {
			if i2 == i1 {
				continue
			}
			for i3 := 1; i3 < 10; i3++ {
				if i3 == i1 || i3 == i2 {
					continue
				}
				for i4 := 1; i4 < 10; i4++ {
					if i4 == i1 || i4 == i2 || i4 == i3 {
						continue
					}
					for i5 := 1; i5 < 10; i5++ {
						if i5 == i1 || i5 == i2 || i5 == i3 || i5 == i4 {
							continue
						}
						//fmt.Printf("%d %d %d %d %d\n", i1, i2, i3, i4, i5)

						sum := 55 + i1 + i2 + i3 + i4 + i5
						if sum%5 != 0 {
							//fmt.Printf("Definitely no good sum\n")
							continue
						}
						sum /= 5
						//fmt.Printf("Sum=%d\n", sum)

						o1 := sum - i1 - i2
						o2 := sum - i2 - i3
						o3 := sum - i3 - i4
						o4 := sum - i4 - i5
						o5 := sum - i5 - i1

						all := []int{i1, i2, i3, i4, i5, o1, o2, o3, o4, o5}
						if len(unique(all)) != 10 {
							//fmt.Printf("Non unique\n")
							continue
						}
						if o1 < 0 || o2 < 0 || o3 < 0 || o4 < 0 || o5 < 0 {
							//fmt.Printf("Negative o\n")
							continue
						}
						if o1 > 10 || o2 > 10 || o3 > 10 || o4 > 10 || o5 > 10 {
							//fmt.Printf("Too big o\n")
							continue
						}
						if o1 > o2 || o1 > o3 || o1 > o4 || o1 > o5 {
							//fmt.Printf("Not starting from smallest\n")
							continue
						}
						cur := fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d", o1, i1, i2, o2, i2, i3, o3, i3, i4, o4, i4, i5, o5, i5, i1)
						//fmt.Println(cur)
						if cur > result {
							result = cur
						}
					}
				}
			}
		}
	}
	return result
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
