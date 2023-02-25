package main

import (
	"fmt"
	"time"
)

//创建有缓存channel
var Queue = make(chan int, 3)

func produceTask() {
	fmt.Printf("len(chan)=%d,cap(chan)=%d \n", len(Queue), cap(Queue))

	for i := 0; i < 4; i++ {
		fmt.Println("Produce task i=", i)
		Queue <- i
	}
	defer func() {
		fmt.Println("processTask completed")
	}()

}

func main() {

	go produceTask()

	time.Sleep(2000 * time.Millisecond)

	for i := 0; i < 4; i++ {
		task := <-Queue
		fmt.Println("Current task is ", task)
	}

	fmt.Println("main completed")
	time.Sleep(2000 * time.Millisecond)
}
