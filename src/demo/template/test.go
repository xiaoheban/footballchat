package main

import (
	"fmt"
	"html/template"
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
	df := "2021-10-01"
	fDate = date.Format(df)
	fmt.Println(fDate)
	return
}
func process1(w http.ResponseWriter, r *http.Request) {
	fuctions := template.FuncMap{"fdate": formatDate}
	t := template.New("templ6.html").Funcs(fuctions)
	t1, _ := t.ParseFiles("templ6.html")
	t1.Execute(w, time.Now()) //格式化日期
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/sample", sampleTemplate)
	http.HandleFunc("/process", process)
	http.HandleFunc("/process1", process1)
	server.ListenAndServe()
}
