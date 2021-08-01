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
	//curl -id "user=wang hui&pwd=123" 127.0.0.1:8080/getBody
	contentLen := r.ContentLength
	body := make([]byte, contentLen)
	r.Body.Read(body)
	fmt.Println(contentLen)
	fmt.Fprintln(w, string(body))
}

/**
处理form表单数据
*/
func processForm(w http.ResponseWriter, r *http.Request) {
	///1.解析Form字段
	///2.解析PostForm字段
	///3.解析MultipartForm字段
	///4.FormValue函数
	///5.PostFormValue函数
	/* 	r.ParseForm()
	   	fmt.Println(r.Form)
	   	fmt.Println(r.Body)
	   	fmt.Fprintln(w, r.Form) */
}

/**
处理json请求体
*/
func processJsonRequest(w http.ResponseWriter, r *http.Request) {

}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/getHeader/", hanleHeaders)
	http.HandleFunc("/getBody/", handleBody)
	http.HandleFunc("/process/", processForm)
	http.HandleFunc("/processJson/", processJsonRequest)
	server.ListenAndServe()
}
