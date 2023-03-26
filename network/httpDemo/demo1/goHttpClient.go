package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	//print base info

	fmt.Println("Header:", resp.Header)
	fmt.Println("Proto:", resp.Proto)
	fmt.Println("Status:", resp.Status)
	fmt.Println("ContentLength:", resp.ContentLength)
	//关闭内容流
	defer resp.Body.Close()

	buf := make([]byte, 4098)
	var content string
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n == 0 {
			fmt.Println("Finished")
			break
		}
		content += string(buf)
	}

	fmt.Printf("|%v|\n", content)
}
