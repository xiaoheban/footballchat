package main

import (
	"html/template"
	"net/http"
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
	t, _ := template.ParseFiles("templ2.html")
	arrs := []string{"jim", "tom", "bob", "steven"}
	t.Execute(w, arrs)
	//设置工作
	//包含工作
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/sample", sampleTemplate)
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
