package httpdemo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleAddUser(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/addUser/", handleAddUser)
	//创建recorder
	writer := httptest.NewRecorder()
	userJson := strings.NewReader(`"id":"1","name":"json"`)
	request, _ := http.NewRequest("POST", "/addUser/", userJson)
	mux.ServeHTTP(writer, request)
	//获取请求
}

func TestHandleGetUser(t *testing.T) {

}
