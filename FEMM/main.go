package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go program"))
	})
	err := http.ListenAndServe(":3001", nil)
	if err == nil {
		fmt.Println("Error on server")
	}
}