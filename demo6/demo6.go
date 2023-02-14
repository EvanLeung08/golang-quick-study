package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Println("number[:3]", numbers[:3])
	fmt.Println("number[:3]", numbers[1:3])
	fmt.Println("number[:3]", numbers[:])

	numbers = append(numbers, 28)
	printSlice(numbers)
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	copy(numbers1, numbers)
	printSlice(numbers1)

	var testMap map[string]string
	testMap = make(map[string]string, 10)
	testMap["1"] = "php"
	testMap["2"] = "go"
	testMap["3"] = "java"
	fmt.Println(testMap)

}

func printSlice(x []int) {
	fmt.Println("len=%d,cap=%d,slice=%v", len(x), cap(x), x)

}
