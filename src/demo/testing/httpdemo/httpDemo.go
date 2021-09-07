package httpdemo

import "net/http"

type User struct {
	Id   string
	Name string
}

/**
增加用户
*/
func handleAddUser(w http.ResponseWriter, r *http.Request) {

}

/**
获取用户信息
*/
func handleGetUser(w http.ResponseWriter, r *http.Request) {

}
