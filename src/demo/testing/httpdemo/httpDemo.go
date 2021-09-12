package httpdemo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

/**
增加用户
*/
func handleAddUser(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(600) //获取入参错误
		return
	}
	var user User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		fmt.Println("parser json err ", err)
		w.WriteHeader(601)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("添加用户成功"))
}

/**
获取用户信息
*/
func handleGetUser(w http.ResponseWriter, r *http.Request) {
	strUser := `{"id":"1","name":"json"}`
	w.WriteHeader(200)
	w.Write([]byte(strUser))
}
