package main

import (
	"fmt"
	"strings"
)

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var valleys, count int32
	paths := strings.Split(path, "")
	for _, s := range paths {
		x := count
		if s == "D" {
			count--
		} else if s == "U" {
			count++
		}

		if count == -1 && x == 0 {
			valleys++
		}
	}

	return valleys
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))
	fmt.Println(countingValleys(12, "DDUUDDUDUUUD"))
}
