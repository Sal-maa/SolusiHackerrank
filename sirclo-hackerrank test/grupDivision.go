package main

import (
	"fmt"
	"sort"
)

func groupDivision(levels []int, maxSpread int) int {
	var result int
	sort.SliceStable(levels, func(i, j int) bool {
		return levels[i] < levels[j]
	})
	var index, jarak int
	var isFound bool
	// fmt.Println(levels)
	for i := 0; i < len(levels); {
		index = i + 1 //pengganti j
		isFound = false
		result++
		for !isFound {
			//sampai ujung array atay ga sampai nemu yang lewat spread
			//index = i + 1
			if index < len(levels) {
				jarak = levels[index] - levels[i]
				if jarak > maxSpread {
					isFound = true
				} else {
					index++ //j++
				}
			} else {
				isFound = true
			}
		}
		i = index //loncat ke penggagal maxspread(ke j)
	}
	//fmt.Println(index,isFound,jarak)
	//fmt.Println(arrayGroup)
	return result
}

func main() {
	fmt.Println(groupDivision([]int{4, 1, 2, 5, 3}, 0))
	fmt.Println(groupDivision([]int{4, 8, 1, 7}, 3))
	fmt.Println(groupDivision([]int{4, 4, 4, 4, 4}, 1))
	fmt.Println(groupDivision([]int{1, 1, 4, 7}, 3))
}
