package encode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func decode(fileName string) (post Post, err error) {
	var bytes []byte
	bytes, err = ioutil.ReadFile(fileName)
	err = json.Unmarshal(bytes, &post)
	return
}

/**
将数据写入json文件
*/
func encode(fileName string, post Post) error {
	datas, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, datas, 0644)
}

//未实现方法
func advanceMethod() {
	//TD DO
}

//耗时方法
func methodWithLongTime() {
	fmt.Println("method test long time")
	time.Sleep(10 * time.Second)
}
