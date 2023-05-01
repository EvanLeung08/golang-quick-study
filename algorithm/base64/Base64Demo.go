package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "fdasfdsa2w32131$%^^&"
	//Encode
	dataEnc := b64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println("EncodedData:" + dataEnc)
	//Decode
	decodeString, err := b64.StdEncoding.DecodeString(dataEnc)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("DecodedData:" + string(decodeString))
}
