package httpdemo

import (
	"encoding/json"
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
	//writer.Body = nil
	userJson := strings.NewReader(`"id":"1","name":"json"`)
	request, _ := http.NewRequest("POST", "/addUser/", userJson)
	mux.ServeHTTP(writer, request)
	//处理响应码
	if writer.Code != 200 {
		t.Errorf("Response Code is %v", writer.Code)
	}
}

func TestHandleGetUser(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/getUser/", handleGetUser)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/getUser?userId=1", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Fatalf("Response Code is %v", writer.Code)
	}
	var userInfo User
	err := json.Unmarshal(writer.Body.Bytes(), &userInfo)
	if err != nil {
		t.Fatalf("parser Json Failed %v", err)
	}
	t.Logf("the user is %v", userInfo)
}
