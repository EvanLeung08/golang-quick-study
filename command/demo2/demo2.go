package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

/*
*
Find the absolute path of the executable file
*/
func main() {
	binary, err := exec.LookPath("ps")
	if err != nil {
		panic(err)
	}
	args := []string{"ps", "-ef"}
	environ := os.Environ()
	err1 := syscall.Exec(binary, args, environ)
	if err1 != nil {
		fmt.Println("err=", err1)
	}
}
