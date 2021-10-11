package footballchat

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("this is foot ball chat")
	mux := http.NewServeMux()

	//设置路由
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", errors)
	mux.HandleFunc("/login", login)

	server := http.Server{
		Addr:    "0.0.0.0:8090",
		Handler: mux,
	}
	server.ListenAndServe()
}
