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

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/sample", sampleTemplate)
	server.ListenAndServe()
}
