package main

import "fmt"

func demo() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

func main() {
	demo()
	recoverDemo(11)
}

func recoverDemo(i int) {
	var arr [10]int
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}

	}()
	arr[i] = 10

}
