package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func write2JsonFile() {
	post := Post{
		Id:      1,
		Author:  "太白",
		Content: "飞流直下三千尺",
		Comments: []Comment{
			{11, "DuFu", "太白真棒"},
			{12, "白居易", "太白真棒+1"},
			{13, "东坡", "太白真棒+2"},
		},
	}
	//	bytes, err := json.Marshal(&post)
	bytes, err := json.MarshalIndent(&post, "", "\t\t")
	if err == nil {
		fmt.Println("wirte json")
		ioutil.WriteFile("test.json", bytes, 0644)
	} else {
		fmt.Println("err", err)
	}
}
func readFromJsonFile() {

}

func main() {
	write2JsonFile()
}
