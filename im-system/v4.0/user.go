package main

import "net"

type User struct {
	Name       string
	IPAddr     string
	Channel    chan string
	Connection net.Conn
	server     *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:       userAddr,
		IPAddr:     userAddr,
		Channel:    make(chan string),
		Connection: conn,
		server:     server,
	}
	//启动消息监听
	go user.ListenMessage()
	return user
}

func (this *User) Online() {
	//把用户加入OnlineMap
	this.server.MapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.MapLock.Unlock()
	//广播用户上线通知
	this.server.Broadcast(this, "已上线！")

}

func (this *User) Offline() {
	this.server.Broadcast(this, "已下线")
	//移除用户
	this.server.OnlineMap[this.Name] = nil
}

func (this *User) processMessage(msg string) {
	//广播通知所有用户
	this.server.Broadcast(this, msg)
}

//监听当前用户channel，有消息就发给客户端
func (this *User) ListenMessage() {
	for {
		message := <-this.Channel

		this.Connection.Write([]byte(message + "\n"))

	}
}
