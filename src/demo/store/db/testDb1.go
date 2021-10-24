package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //导入数据库,加上下划线表示仅仅调用其init方法完成导入和驱动初始化
)

type MyPost struct {
	Id string
}

var db *sql.DB //全局使用

//连接数据库
func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://test_db:test1234@119.45.54.127:5432/test?ssmode=disable")
	if err != nil {
		fmt.Println("打开数据库失败", err)
	} else {
		fmt.Println("连接数据库成功")
	}
}

//新增帖子

//删除帖子
//查询帖子
func main() {

}
