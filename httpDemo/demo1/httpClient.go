package main

import (
	"fmt"
	"net"
	"os"
)

func errorFunc1(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

// 模拟浏览器
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	errorFunc1(err, "net.dial err:")
	defer conn.Close()

	httpRequest := "GET /test1 HTTP/1.1\r\nHost: 127.0.0.1:8000\r\n\r\n"
	conn.Write([]byte(httpRequest))

	//读响应信息
	buf := make([]byte, 4098)
	n, err := conn.Read(buf)
	errorFunc1(err, "conn.Read err:")

	if n == 0 {
		return
	}

	fmt.Printf("|%s|\n", buf[:n])
}
