package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
   <title>Go语言标准库文档中文版 | Golang中国</title>
   <meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
   <meta charset="utf-8">
   <link rel="shortcut icon" href="/static/img/go.ico">
</head>
       <title>标题</title>
        <div>过年来吃鸡啊</div>
        <div>hello regexp</div>
        <div>你在吗？</div>
       <div>
         2块钱啥时候还？
         过了年再说吧！
         刚买了车，没钱。。。
       </div>
        <body>呵呵</body>
 
<frameset cols="15,85">
   <noframes>
   </noframes>
</frameset>
</html>`

	regex := regexp.MustCompile("<div>(.*?)</div>")
	array := regex.FindAllStringSubmatch(str, -1)

	for i := 0; i < len(array); i++ {
		one := array[i]
		for j := 0; j < len(one); j++ {
			fmt.Printf("%d:%v\n", j, one[j])
		}
	}
	fmt.Println(array)
	for _, str := range array {
		fmt.Println("first search:", str[1])
	}

	regex1 := regexp.MustCompile("<div>(?s:(.*?))</div>")
	array1 := regex1.FindAllStringSubmatch(str, -1)

	for _, str1 := range array1 {
		fmt.Println("second search:", str1[1])
	}

}
