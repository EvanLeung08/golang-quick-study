package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 对外提供一个方法创建服务器实例
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接成功", conn)

}

// 启动服务器
func (this *Server) Start() {
	//建立链接
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("accept error is :", err)
		return
	}
	//关闭链接
	defer listener.Close()
	//接收数据
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listener accepted error is :", err)
			continue
		}
		//处理信息
		go this.Handler(conn)
	}

}
