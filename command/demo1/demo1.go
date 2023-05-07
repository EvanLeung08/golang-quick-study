package main

import (
	"fmt"
	"io"
	"os/exec"
)

/*
*
Execute external command
*/
func main() {

	dateCmd := exec.Command("date")
	output, err := dateCmd.Output()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("> date")
	fmt.Println(string(output))
	_, err1 := exec.Command("date", "-x").Output()
	if err1 != nil {
		switch e := err1.(type) {
		case *exec.Error:
			fmt.Println("Failed to execute command:", err1)
		case *exec.ExitError:
			fmt.Println("Command exit , rc=", e.ExitCode())
		default:
			fmt.Println("Unknown error:", err1)
		}
	}

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
