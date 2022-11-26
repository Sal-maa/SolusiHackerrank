package main

import "fmt"

func viralAdvertising(n int32) int32 {
	// Write your code here
	var liked, commulative int32
	shared := int32(5)
	for i := int32(1); i <= n; i++ {
		liked = shared / 2
		commulative = commulative + liked
		shared = liked * 3
	}

	return commulative
}

func main() {
	fmt.Println(viralAdvertising(34))
	fmt.Println(viralAdvertising(50))
}
