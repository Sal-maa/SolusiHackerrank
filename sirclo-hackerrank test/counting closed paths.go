package main

import (
	"fmt"
	"strconv"
)

func closedPaths(number int32) int32 {
	// Write your code here
	sNumber := strconv.Itoa(int(number))
	count := 0

	for i := 0; i < len(sNumber); i++ {
		a, _ := strconv.Atoi(string(sNumber[i]))
		if a == 0 || a == 4 || a == 6 || a == 9 {
			count++
		} else if a == 8 {
			count += 2
		}
	}
	return int32(count)
}

func main() {
	fmt.Println(closedPaths(649578))
}
