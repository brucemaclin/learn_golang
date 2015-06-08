package main

import (
	"fmt"
	"math/rand"
)

func bubberSort(data []int, length int) {
	if length < 0 || length < 2 {
		return
	}
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}

	}
}

func main() {
	var data [10]int
	for i := 0; i < 10; i++ {
		data[i] = rand.Intn(100) + 1
	}
	bubberSort(data[:10], 10)
	fmt.Println(data)
}
