package main

import "fmt"

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	if p > s {
		return 0
	} else if p == s {
		return 1
	}

	g := int32(0)
	for {
		s = s - p
		if s < 0 {
			break
		}

		g++

		p = p - d
		if p <= m {
			p = m
		}
	}

	return g
}

func main() {
	fmt.Println(howManyGames(20, 3, 6, 80))
	fmt.Println(howManyGames(1, 100, 1, 9777))
	fmt.Println(howManyGames(100, 3, 100, 6234))
	fmt.Println(howManyGames(73, 72, 44, 9163))
	fmt.Println(howManyGames(73, 51, 43, 2873))
	fmt.Println(howManyGames(89, 91, 26, 840))
	fmt.Println(howManyGames(39, 99, 12, 9289))
	fmt.Println(howManyGames(58, 50, 9, 6657))
}
