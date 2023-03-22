package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/eshare", handleCallback)

	http.ListenAndServe("127.0.0.1:8000", nil)
}

func handleCallback(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("This is a web server!"))

	fmt.Println("Header", request.Header)
	fmt.Println("Host", request.Host)
	fmt.Println("RemoteAddr", request.RemoteAddr)
	fmt.Println("Body", request.Body)

}
