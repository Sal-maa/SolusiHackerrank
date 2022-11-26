package main

import (
	"fmt"
	"strings"
)

func camelcase(s string) int32 {
	// Write your code here
	snakecase := strings.ToLower(s)
	mapString := make(map[int]rune)
	for idx, str := range s {
		mapString[idx] = str
	}
	var count int32
	for idx, strCheck := range snakecase {
		if mapString[idx] != strCheck {
			count++
		}
	}
	return count + 1
}

func main() {
	fmt.Println(camelcase("abCd"))
	fmt.Println(camelcase("abcd"))
	fmt.Println(camelcase("abc"))
	fmt.Println(camelcase("ab"))
	fmt.Println(camelcase("singleword"))
	fmt.Println(camelcase("saveChangesInTheEditor"))
}
