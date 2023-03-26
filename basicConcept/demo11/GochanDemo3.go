package main

import "fmt"

func main() {
	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()
	x, y := 1, 1

	for {
		select {
		case channel <- x:
			x = y
			y = x + y
		case <-quit:
			return
		}
	}

	fmt.Println("server 线程结束")
}
