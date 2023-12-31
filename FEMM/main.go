package main

import (
	"fmt"
	"net/http"
)
func handleHello (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from go program"))
}
func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)

	err := http.ListenAndServe(":3001", server)
	if err == nil {
		fmt.Println("Error on server")
	}
}