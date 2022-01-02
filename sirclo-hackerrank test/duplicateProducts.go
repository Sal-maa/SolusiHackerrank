package main

import "fmt"

func duplicateProduct(name []string, price []int32, waight []int32) int32 {
	var count int32

	strSlice := []string{}
	for i := 0; i < len(name); i++ {
		str := ""
		str = str + name[i] + "," + fmt.Sprint(price[i]) + "," + fmt.Sprint(waight[i])
		strSlice = append(strSlice, str)
	}

	strMap := map[string]int32{}
	for _, v := range strSlice {
		strMap[v] += 1
	}

	for _, v := range strMap {
		count += v
	}

	hasil := count - int32(len(strMap))
	return hasil
}

func main() {
	fmt.Println(duplicateProduct([]string{"ball", "box", "ball", "ball", "box"}, []int32{2, 2, 2, 2, 2}, []int32{1, 2, 1, 1, 3}))
}
