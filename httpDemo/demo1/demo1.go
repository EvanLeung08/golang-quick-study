package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errorFunc(err, "net.Listen err:")
	defer listener.Close()
	conn, err := listener.Accept()
	errorFunc(err, "listener.Accept err:")
	b := make([]byte, 1024)
	n, err := conn.Read(b)
	errorFunc(err, "net.dial err:")
	fmt.Printf("|%s|\n", string(b[:n]))
	select {}
}

func errorFunc(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
