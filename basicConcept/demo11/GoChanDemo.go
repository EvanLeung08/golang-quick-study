package main

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func processFile() {
	fmt.Println("Start to process file")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("File Process Done")
	//发送666到channel
	channel <- 666
}

func main() {

	go processFile()
	//从Channel接受数据
	status := <-channel
	fmt.Printf("Received status=>%d", status)

}
