package main

import "fmt"

func main() {
	//要删除slice中间的某个元素并保存原有的元素顺序
	//{5, 6, 7, 8, 9} ——> {5, 6, 8, 9}

	slice := []int{5, 6, 7, 8, 9}
	slice = removeSlice(slice, 3)

	fmt.Println(slice)

}

func removeSlice(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
