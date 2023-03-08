package main

import (
	"flag"
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

var serverIp string

var serverPort int

func init() {
	flag.StringVar(&serverIp, "serverIp", "127.0.0.1", "服务器默认值(127.0.0.1)")
	flag.IntVar(&serverPort, "serverPort", 8888, "服务器默认端口8888")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>>服务器链接失败")
		return
	}

	fmt.Println(">>>>>>服务器链接成功")
	select {}
}
