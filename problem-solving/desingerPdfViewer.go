package main

import (
	"fmt"
	"strings"
)

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	alphabeth := "abcdefghijklmnopqrstuvwxyz"

	charIdx := make([]int32, 0)
	for _, w := range word {
		idx := int32(strings.Index(alphabeth, string(w)))
		charIdx = append(charIdx, h[idx])
	}

	max := charIdx[0]
	for _, idx := range charIdx {
		if idx >= max {
			max = idx
		}
	}
	return (int32(len(charIdx)) * max)
}

func main() {
	fmt.Println(designerPdfViewer([]int32{6, 5, 7, 3, 6, 7, 3, 4, 4, 2, 3, 7, 1, 3, 7, 4, 6, 1, 2, 4, 3, 3, 1, 1, 3, 5}, "zbkkfhwplj"))
	fmt.Println(designerPdfViewer([]int32{6, 3, 4, 4, 6, 4, 5, 3, 4, 3, 6, 5, 4, 6, 7, 1, 3, 4, 2, 5, 6, 1, 5, 1, 7, 2}, "nrdyluacvr"))
}
