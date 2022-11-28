package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findDigits(n int32) int32 {
	// Write your code here
	var count int32
	nSplit := strings.Split(strconv.Itoa(int(n)), "")
	for _, s := range nSplit {
		nInt, _ := strconv.Atoi(s)
		if nInt != 0 && n%int32(nInt) == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(findDigits(114108089))
	fmt.Println(findDigits(110110015))
}
