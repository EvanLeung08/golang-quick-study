package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	id   string
	age  int
}

func (this *User) GetName() {
	fmt.Println(this.name)
}

func printType(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg).Name())
}

func main() {

	var num float32

	printType(num)

	user := User{"haha", "123", 12}

	printClassDetails(user)

}

func printClassDetails(input interface{}) {
	//反射获取对象类型
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is ", inputType.Name())

	//反射获取对象值
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is ", inputValue)

	//获取该类所有字段 和 对应的值
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i)
		fmt.Printf("%s:%v is %v \n", field.Name, field.Type, value)
	}
	//获取该类所有方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s ,%v \n", m.Name, m.Type)
	}
}
