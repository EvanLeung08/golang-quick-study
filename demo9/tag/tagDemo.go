package main

import (
	"fmt"
	"reflect"
)

type Form struct {
	Name   string `info:"name" doc:"我的名字"`
	Sexual string `info:"sexual" doc:"性别"`
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()
	for i := 0; i < t.NumField(); i++ {
		info := t.Field(i).Tag.Get("info")
		doc := t.Field(i).Tag.Get("doc")

		fmt.Printf("info=%s,doc=%s \n", info, doc)
	}

}

func main() {

	form := Form{"小花花", "女"}
	findTag(&form)
}
