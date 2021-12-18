package main

import (
	"fmt"
	"strings"
)

func getMinimumDifference(a []string, b []string) []int32 {
	// Write your code here
	hasil := []int32{}

	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			hasil = append(hasil, -1)
		}
		if len(a[i]) == len(b[i]) {
			aArray := strings.Split(a[i], "")
			bArray := strings.Split(b[i], "")

			label := map[string]bool{}

			for _, v := range aArray {
				label[v] = true
			}
			// fmt.Println(label)
			for _, v := range bArray {
				label[v] = false
			}

			// fmt.Println(label)

			diff := int32(0)
			for _, v := range label {
				if v == true {
					diff += 1
				}
			}
			hasil = append(hasil, diff)
		}
	}
	return hasil
}

func main() {
	fmt.Println(getMinimumDifference([]string{"jk", "abb", "mn", "abc", "a"}, []string{"kj", "bbc", "op", "def", "bb"}))

	// kata := "saye"

	// huruf := strings.Split(kata, "")
	// fmt.Println(huruf)
	// label := map[string]int{}
	// for _, v := range huruf {
	// 	label[v] = 1
	// }
	// fmt.Println(label)
}
