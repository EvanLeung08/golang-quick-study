package main

import "fmt"

type Store struct {
	name    string
	address string
}

func checkType(arg interface{}) {

	fmt.Println(arg)

	value, ok := arg.(Store)

	if !ok {
		fmt.Println("arg is not Store")
	} else {
		fmt.Println("arg is store , value is ", value)
		fmt.Printf("arg type is %T \n", value)
	}

}

func main() {
	store := Store{"阿里巴巴", "杭州"}
	checkType(store)

}
