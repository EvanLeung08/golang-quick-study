package main

import (
	"fmt"
	"golang-quick-study/basicConcept/demo4/lib1"
	"golang-quick-study/basicConcept/demo4/lib2"
	_ "golang-quick-study/demo4/lib1"
)

func init() {
	fmt.Println("libmain init")
}

func main() {
	lib1.Lib1test()
	lib2.Lib2test()
	var x string = "haha"
	var y string = "baba"
	swap(&x, &y)
	fmt.Println(x, y)

	var a string = "haha"
	var b string = "baba"

	swapValue(a, b)
	fmt.Println(a, b)

	a, b = swapValue(a, b)
	fmt.Println(a, b)

}

func swap(x, y *string) {
	var tmp = *x
	*x = *y
	*y = tmp
}

func swapValue(x, y string) (string, string) {
	var tmp = x
	x = y
	y = tmp
	return x, y
}
