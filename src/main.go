package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("这是一个足球迷论坛网站")
	mux := http.NewServeMux()
	mux.HandleFunc()
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
