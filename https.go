package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "https server of golang")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":8086", "server.pem", "server.key", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}