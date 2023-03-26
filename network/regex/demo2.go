package main

import (
	"fmt"
	"regexp"
)

func main() {
	test := "1.2,3425,22.3,22.3333,231,012"

	regex := regexp.MustCompile("[0-9]+\\.[0-9]+")

	submatch := regex.FindAllStringSubmatch(test, -1)

	fmt.Printf("所有匹配的字符是:%v", submatch)
}
