package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	maxA := int32(0)
	for _, v := range a {
		if v >= maxA {
			maxA = v
		}
	}
	minB := b[0]
	for _, v := range b {
		if v <= minB {
			minB = v
		}
	}

	hasil := []int32{}
	for i := maxA; i <= minB; i++ {
		faktorA := 0
		faktorB := 0
		for _, x := range b {
			if x%i == 0 {
				faktorB++
			}
		}
		for _, y := range a {
			if i%y == 0 {
				faktorA++
			}
		}
		if faktorA == len(a) && faktorB == len(b) {
			hasil = append(hasil, i)
		}
	}
	return int32(len(hasil))
}

func main() {
	a := []int32{3, 4}
	b := []int32{24, 48}
	fmt.Println("==============BetweenTwoSets=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("\t  b = ", b)
	fmt.Println("Output\t: ", getTotalX(a, b))
}
