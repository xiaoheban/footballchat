package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	/* r.ParseForm()
	fmt.Fprintln(w, r.Form)  */
	//enctype="application/x-www-form-urlencoded"
	//output:map[jsj:[89] pwd:[sss] user:[wang hui]] wang来自form表单 hui来自url拼接的参数
	///2.解析PostForm字段
	/* r.ParseForm()
	fmt.Fprintln(w, r.PostForm) */
	///3.解析MultipartForm字段
	///4.FormValue函数
	//fmt.Fprintln(w, r.FormValue("user")) //wang
	///5.PostFormValue函数
	//fmt.Fprintln(w, r.PostFormValue("user")) //wang
	/* 	r.ParseForm()
	   	fmt.Println(r.Form)
	   	fmt.Println(r.Body)
	   	fmt.Fprintln(w, r.Form) */
}

/**
form表单上传文件
*/
func handleFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request here")
	//r.ParseMultipartForm(2048)
	//fmt.Fprintln(w, r.MultipartForm)//&{map[file_name:[ss]] map[apk:[0xc000040050]]}
	/* 	fmt.Println("fileParams", r.MultipartForm.File["apk"])
	   	fileHandler := r.MultipartForm.File["apk"][0]
	   	file, err := fileHandler.Open()
	   	if err == nil {
	   		contents, err := ioutil.ReadAll(file)
	   		if err == nil {
	   			fmt.Fprintln(w, string(contents))
	   		}
	   	} */
	//	//可以使用简化的formFile("apk")来简化操作，这样只会返回第一个文件的句柄
	file, _, err := r.FormFile("apk")
	if err == nil {
		contents, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(contents))
		}
	}
}

/**
处理json请求体
*/
type LoginInfo struct {
	Id       string
	Account  string
	Password string
}

func processJsonRequest(w http.ResponseWriter, r *http.Request) {
	var loginInfo LoginInfo
	err := json.NewDecoder(r.Body).Decode(&loginInfo)
	if err != nil {
		fmt.Fprintln(w, loginInfo)
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/getHeader/", hanleHeaders)
	http.HandleFunc("/getBody", handleBody)
	http.HandleFunc("/process", processForm)
	http.HandleFunc("/handleFile", handleFile)
	http.HandleFunc("/processJson", processJsonRequest)
	server.ListenAndServe()
}
