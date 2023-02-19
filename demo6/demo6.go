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

	testMap2 := map[string]string{
		"test1": "1",
		"test2": "2",
		"test3": "3",
	}
	fmt.Println(testMap2)

	testMap3 := make(map[string]map[string]string)
	testMap3["item1"] = make(map[string]string, 2)
	testMap3["item1"]["id"] = "1"
	testMap3["item1"]["desc"] = "Go语言很棒"
	testMap3["item2"] = make(map[string]string, 2)
	testMap3["item2"]["id"] = "2"
	testMap3["item2"]["desc"] = "java语言很棒"
	fmt.Println(testMap3)

}

func printSlice(x []int) {
	fmt.Println("len=%d,cap=%d,slice=%v", len(x), cap(x), x)

}
