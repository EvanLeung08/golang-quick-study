package main

import "fmt"

var x, y int
var (
	a int
	b bool
)

var c, d int = 10, 20
var e, f = 12, "evan liang"

func main() {
	g, h := 20, "只能用于方法内"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	//舍弃某些值
	_, i := 7, 5

	fmt.Printf("i=%d\n", i)
}
