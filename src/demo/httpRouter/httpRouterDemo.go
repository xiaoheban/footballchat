/**
使用httpRouter这个多路复用器
*/
package main

/**
httpRouter和原生的对比，获取参数更加方便了
*/
import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hanleEngland(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "manchister city is a gread club! %v", p)
}
func main() {
	mux := httprouter.New()
	mux.GET("/england_club/", hanleEngland)

	server := http.Server{
		Addr:    "127.0.0.1:8090",
		Handler: mux,
	}
	server.ListenAndServe()
}
