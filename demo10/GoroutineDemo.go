package main

import (
	"fmt"
	"time"
)

func doTask() {
	i := 0
	for {
		i++
		fmt.Printf("taskId = %d \n", i)
	}

}

func main() {

	/* go doTask()

	i := 0
	for {
		i++
		fmt.Printf("mainTask = %d \n", i)

	}*/

	go func() {

		defer func() {
			fmt.Println("A")
		}()

		fmt.Println("B")

	}()
	fmt.Println("C")
	time.Sleep(1000)
}
