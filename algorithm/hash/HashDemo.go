package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Sha256 test"

	hash := sha256.New()
	hash.Write([]byte(s))
	hashValue := hash.Sum(nil)

	fmt.Println("original text:" + s)
	fmt.Printf("hash result:%x\n", hashValue)
}
