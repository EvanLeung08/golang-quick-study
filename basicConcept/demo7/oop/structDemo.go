package main

import "fmt"

type Good struct {
	title  string
	author string
}

func changeGood1(Good Good) {
	Good.title = "haha1"
	Good.author = "Evan1"
}

func changeGood2(Good *Good) {
	Good.title = "haha2"
	Good.author = "Evan2"
}

func main() {
	var Good Good
	Good.title = "haha"
	Good.author = "Evan"
	fmt.Println(Good)
	changeGood1(Good)
	fmt.Println(Good)

	changeGood2(&Good)
	fmt.Println(Good)
}
