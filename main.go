package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int
	for i := 0; i < 20; i++ {
		numbers = append(numbers, rand.Intn(100))
	}
	fmt.Println(numbers)
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j+1], numbers[j] = numbers[j], numbers[j+1]
			}
		}
	}
	fmt.Println(numbers)
}
