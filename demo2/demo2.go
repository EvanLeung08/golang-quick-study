package main

import "fmt"

func main() {
	//使用默认值
	var a int
	fmt.Printf("a=%d\n", a)

	//使用赋值
	var b int = 10
	fmt.Printf("b=%d\n", b)

	//省略后面类型，自动匹配
	var c = 12
	fmt.Printf("%d\n", c)

	//省略声明
	d := 20
	fmt.Printf("%d\n", d)
}
