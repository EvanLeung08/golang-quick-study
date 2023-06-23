package main

import (
	"fmt"
	"regexp"
)

func main() {
	test := "1@@,163.com,1234@163.com,fdsaf@qq.com"
	//编译程正则对象
	regex := regexp.MustCompile("[0-9a-zA-Z]+@[0-9a-zA-Z]+\\.com")
	//通过正则对象取匹配字符
	array := regex.FindAllStringSubmatch(test, -1)
	fmt.Printf("匹配到的字符:%v", array)

}
