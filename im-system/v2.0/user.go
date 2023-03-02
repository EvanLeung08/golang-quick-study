package main

import "net"

type User struct {
	Name       string
	IPAddr     string
	Channel    chan string
	Connection net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:       userAddr,
		IPAddr:     userAddr,
		Channel:    make(chan string),
		Connection: conn,
	}
	//启动消息监听
	go user.ListenMessage()
	return user
}

//监听当前用户channel，有消息就发给客户端
func (this *User) ListenMessage() {
	for {
		message := <-this.Channel

		this.Connection.Write([]byte(message + "\n"))

	}
}
