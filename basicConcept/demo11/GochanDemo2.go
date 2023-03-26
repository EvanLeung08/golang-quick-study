package main

import "fmt"

func main() {

	channel := make(chan int, 3)
	go func() {

		for i := 0; i < 4; i++ {
			fmt.Println("Generate message ", i)
			channel <- i
		}
		defer fmt.Println("Go子线程结束")
		close(channel)
	}()

	for i := 0; i < 3; i++ {

		/*	if data, ok := <-channel; ok {
				fmt.Println("输出:", data)
			} else {
				break
			}*/
		for data := range channel {
			fmt.Println("输出:", data)
		}
	}
	fmt.Println("main结束")
}
