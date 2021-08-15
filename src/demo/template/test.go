package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

/**
简单的模板引擎
*/
func sampleTemplate(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("temp.html")
	template.Execute(w, "hello world")
}
func process(w http.ResponseWriter, r *http.Request) {
	//条件动作
	/* 	t, _ := template.ParseFiles("templ1.html")
	   	rand.Seed(time.Now().Unix())
	   	t.Execute(w, rand.Intn(10) > 5) */
	//迭代动作
	/* 	t, _ := template.ParseFiles("templ2.html")
	   	arrs := []string{"jim", "tom", "bob", "steven"}
	   	t.Execute(w, arrs) */
	//设置工作
	/* 	t, _ := template.ParseFiles("templ3.html")
	   	t.Execute(w, "hello") */
	//包含工作
	t, _ := template.ParseFiles("templ4.html", "templ5.html")
	t.Execute(w, "hello go")
}
func formatDate(date time.Time) (fDate string) {
	df := "2006-01-01" //固定
	fDate = date.Format(df)
	fmt.Println(fDate)
	return
}

/**
自定义函数
*/
func process1(w http.ResponseWriter, r *http.Request) {
	fuctions := template.FuncMap{"fdate": formatDate}
	t := template.New("templ6.html").Funcs(fuctions)
	t1, _ := t.ParseFiles("templ6.html")
	t1.Execute(w, time.Now()) //格式化日期
}

/**
模板嵌套
*/
func process2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templ7.html")
	if err == nil {
		t.ExecuteTemplate(w, "layout", "") //注意这里第二个参数不再是数据，二是模板名
	}
	fmt.Println(err)
}

/**
利用随机数模仿不同文件定义相同的模板 content
*/
func process3(w http.ResponseWriter, r *http.Request) {
	var t template.Template
	rand.Seed(time.Now().Unix())
	if rand.Intn(100) > 20 {
		t = *template.Must(template.ParseFiles("templ7.html", "content-red.html"))
	} else {
		t = *template.Must(template.ParseFiles("templ7.html", "content-blue.html"))
	}
	t.ExecuteTemplate(w, "layout", "")
}

/**
定义默认模板
*/
func process4(w http.ResponseWriter, r *http.Request) {
	var t template.Template
	rand.Seed(time.Now().Unix())
	if rand.Intn(100) > 20 {
		t = *template.Must(template.ParseFiles("templ9.html", "content-red.html"))
	} else {
		t = *template.Must(template.ParseFiles("templ9.html"))
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/sample", sampleTemplate)
	http.HandleFunc("/process", process)
	http.HandleFunc("/process1", process1)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/process3", process3)
	http.HandleFunc("/process4", process4)

	server.ListenAndServe()
}
