package main

import (
	"fmt"
	"net/http"
)

func hanleHeaders(w http.ResponseWriter, r *http.Request) {
	//h := r.Header //取得所有首部
	//h := r.Header["Accept-Encoding"] //获取某个首部 output:[gzip, deflate, br] 数组
	h := r.Header.Get("Accept-Encoding") //获取某个首部 是,分割的 gzip, deflate, br
	fmt.Fprintln(w, h)
}
func handleBody(w http.ResponseWriter, r *http.Request) {
	contentLen := r.ContentLength
	body := make([]byte, contentLen)
	r.Body.Read(body)
	fmt.Println(contentLen)
	fmt.Fprintln(w, string(body))
}

/**
解析form表单数据
*/
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.Body)
	fmt.Fprintln(w, r.Form)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/getHeader/", hanleHeaders)
	http.HandleFunc("/getBody/", handleBody)
	http.HandleFunc("/process/", process)
	server.ListenAndServe()
}
