package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var start int
	var end int

	fmt.Println("请输入开始页:")
	fmt.Scan(&start)
	fmt.Println("请输入结束页:")
	fmt.Scan(&end)

	//开始爬取百度数据
	working(start, end)
}

func working(start int, end int) {
	fmt.Printf("正在爬取%d到%d页", start, end)

	for i := start; i <= end; i++ {
		var result string
		//访问百度贴吧
		url := "https://tieba.baidu.com/f?kw=%E5%8C%97%E4%BA%AC%E7%90%86%E5%B7%A5%E5%A4%A7%E5%AD%A6%E7%8F%A0%E6%B5%B7%E5%AD%A6%E9%99%A2&ie=utf-8&pn="
		result, err := HttpGet(url + strconv.Itoa((i-1)*50))
		fmt.Printf("正在爬%d页,url=%s\n", i, url)
		if err != nil {
			fmt.Println("HttpGet error", err)
			continue
		}
		fmt.Println(result)

		//创建文件
		fileName := strconv.Itoa(i) + ".html"
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println("os.Create err", err)
			continue
		}
		//导出数据
		f.Write([]byte(result))
		//关闭文件流
		f.Close()
	}
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http.Get err:", err)
		err = err1
	}
	defer resp.Body.Close()
	buf := make([]byte, 4098)
	for {
		n, err2 := resp.Body.Read(buf)

		if n == 0 {
			fmt.Println("Finished")
			break
		}

		if err2 != nil && err2 != io.EOF {
			err = err2
			break
		}

		result += string(buf[:n])
	}
	return result, err
}
