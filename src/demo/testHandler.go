/**
演示如何使用默认的多路复用器绑定多个处理器
*/
package main

import (
	"fmt"
	"net/http"
)

type ChinaHander struct{}
type USAHandler struct{}

func (china ChinaHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello china")
}
func (usa USAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

///处理器函数
func handleJapa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hellpo japa")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/china/", ChinaHander{}) // /china ok /china/aa ok
	http.Handle("/usa", USAHandler{})     // /usa ok /usa/aaa not ok

	http.HandleFunc("/japa/", handleJapa)

	server.ListenAndServe()

}
