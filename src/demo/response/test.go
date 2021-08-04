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
func handleWrite(w http.ResponseWriter, r *http.Request) {
	strResponse := `<html><header><header><body><div>你好</div></body></html>`
	w.Write([]byte(strResponse))
}
func handleWriteHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "api not implemented!!")
}
func handleHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://baidu.com")
	w.WriteHeader(302)
}

type Post struct {
	Author  string   //作者
	Threads []string //帖子
}

func handleResponseJson(w http.ResponseWriter, r *http.Request) {
	post := Post{
		Author:  "jim",
		Threads: []string{"曼城8球大胜曼联", "富登上演助攻帽子戏法", "阿奎罗前来观战", "新人斯特林表现抢眼"},
	}
	respData, err := json.Marshal(post)
	fmt.Println(string(respData))
	if err == nil {
		//	w.Header().Set("Content-Type", "applicaiton/json") //不写不会自动识别
		w.Write(respData)
	}
}
func handleSetCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:  "c1",
		Value: "12",
	}
	c2 := http.Cookie{
		Name:     "c2",
		Value:    "34",
		HttpOnly: true,
	}
	c3 := http.Cookie{
		Name:     "c3",
		Value:    "565",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c3)
}
func getCookie(w http.ResponseWriter, r *http.Request) {
	//cookie := r.Header["Cookie"] //返回切片
	//	c1, _ := r.Cookie("c1") //返回单个cookie
	cookies := r.Cookies() //返回切片
	fmt.Fprintln(w, cookies)
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
	http.HandleFunc("/handleWrite", handleWrite)
	http.HandleFunc("/handleWriteHeader", handleWriteHeader)
	http.HandleFunc("/handleHeader", handleHeader)
	http.HandleFunc("/handleResponseJson", handleResponseJson)
	http.HandleFunc("/handleSetCookie", handleSetCookie)
	http.HandleFunc("/getCookie", getCookie)

	server.ListenAndServe()
}
