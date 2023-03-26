package lib1

import (
	"fmt"
	_ "golang-quick-study/demo4/lib2"
)

func init() {
	fmt.Println("lib1 init")
}

func Lib1test() {
	fmt.Println("lib1test")
}
