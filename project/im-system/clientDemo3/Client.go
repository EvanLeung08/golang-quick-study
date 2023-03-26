package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	Name   string
	IPAddr string
	Port   int
	C      net.Conn
	flag   int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		Name:   serverIp,
		IPAddr: serverIp,
		Port:   serverPort,
		flag:   999,
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

func (this *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改名称")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if 0 <= flag && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>>>菜单输入有误，请重新输入<<<<<<<<<\n")
		return false
	}

}
func (this *Client) Rename() bool {
	fmt.Println("请输入用户名:")
	fmt.Scanln(&this.Name)
	sendMsg := "rename|" + this.Name + "\n"
	_, err := this.C.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("C.Write err：", err)
		return false
	}

	return true
}

func (this *Client) DealResponse() {
	//阻塞舰艇channel的数据
	io.Copy(os.Stdout, this.C)

}

func (this *Client) Run() {
	for this.flag != 0 {
		for this.menu() != true {
		}
		switch this.flag {
		case 1:
			fmt.Println("公聊模式")
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			//更改用户名
			this.Rename()
			break
		}
	}
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>>服务器链接失败")
		return
	}
	go client.DealResponse()

	fmt.Println(">>>>>>服务器链接成功")
	client.Run()
}
