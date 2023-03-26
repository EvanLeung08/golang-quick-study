package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	Message   chan string
	mapLock   sync.RWMutex
}

// 对外提供一个方法创建服务器实例
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接成功:", conn.RemoteAddr().String())
	//创建用户实例
	user := NewUser(conn)
	//把用户加入OnlineMap
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	//广播用户上线通知
	this.Broadcast(user, "已上线！")

	//启动协程监听用户输入
	go func() {
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("Received error :", err)
				return
			}
			if n == 0 {
				this.Broadcast(user, "已下线")
				//移除用户
				this.OnlineMap[user.Name] = nil
				return
			}

			msg := string(buf[:n-1])
			//广播通知所有用户
			this.Broadcast(user, msg)
		}
	}()

	//阻塞当前协程
	select {}

}

func (this *Server) ListenMessage() {
	//持续监听消息通道，如果收到消息则通知所有用户
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		//获取所有用户实例
		for _, user := range this.OnlineMap {
			user.Channel <- msg
		}
		this.mapLock.Unlock()
	}

}

func (this *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.IPAddr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
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
	//开启协程监听
	go this.ListenMessage()
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
