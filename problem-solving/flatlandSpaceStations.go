package main

import (
	"fmt"
	"math"
)

func flatlandSpaceStations(n int32, c []int32) int32 {
	if n <= int32(len(c)) {
		return 0
	}

	jarak := make([]int32, 0)
	for i := int32(0); i < n; i++ {
		min := int32(math.Abs(float64(c[0] - i)))
		for j := int32(1); j < int32(len(c)); j++ {
			res := int32(math.Abs(float64(c[j] - i)))
			if res < min {
				min = res
			}
		}
		jarak = append(jarak, min)
	}

	max := int32(0)
	for _, v := range jarak {
		if v >= max {
			max = v
		}
	}

	return max
}

func main() {
	fmt.Println(flatlandSpaceStations(5, []int32{0, 4}))
	fmt.Println(flatlandSpaceStations(6, []int32{0, 1, 2, 3, 4, 5}))
	fmt.Println(flatlandSpaceStations(100, []int32{93, 41, 91, 61, 30, 6, 25, 90, 97}))
	fmt.Println(flatlandSpaceStations(97, []int32{71, 72, 0, 28, 79, 4, 73, 56, 78, 83, 95, 26, 49, 27, 91, 77, 16, 20, 87, 63, 33, 36, 7, 23, 92, 74, 80, 68, 57, 62, 52, 84, 50, 13, 69, 39, 90, 55, 17, 64, 81, 22, 88, 8, 46, 85, 44, 96, 35, 47, 89, 93, 11, 75, 38, 29, 86, 19, 70, 42, 30, 58, 82, 76, 48, 54, 9}))
	fmt.Println(flatlandSpaceStations(5, []int32{0, 4}))
}
