package main

import (
	"fmt"
	"net"
)

type Client struct {
	Name   string
	IPAddr string
	Port   int
	C      net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		Name:   serverIp,
		IPAddr: serverIp,
		Port:   serverPort,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error", err)
		return nil
	}
	client.C = conn

	return client
}

func main() {
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>>>>>>服务器链接失败")
		return
	}

	fmt.Println(">>>>>>服务器链接成功")
	select {}
}
