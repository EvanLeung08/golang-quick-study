package main

import "net/http"

func main() {
	http.HandleFunc("/test", callback)

	http.ListenAndServe("127.0.0.1:8000", nil)
}

func callback(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello HTTP"))
}
