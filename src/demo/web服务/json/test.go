package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	//方法1
	//	bytes, err := json.Marshal(&post)
	bytes, err := json.MarshalIndent(&post, "", "\t\t")
	if err == nil {
		fmt.Println("wirte json")
		ioutil.WriteFile("test.json", bytes, 0644)
	} else {
		fmt.Println("err", err)
	}
	// 方法2
	file, _ := os.Create("test2.json")
	defer file.Close()
	jsonEncoder := json.NewEncoder(file)
	if jsonEncoder != nil {
		jsonEncoder.Encode(&post)
	}
}
func readFromJsonFile() {
	//方法一，先把数据读取到内存，再转成json
	bytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	var post1 Post
	err = json.Unmarshal(bytes, &post1)
	if err != nil {
		panic(err)
	}
	fmt.Println("get post from file use Unmarshal", post1)
	//方法2,直接从文件读j'son
	file, err := os.Open("test2.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decorder := json.NewDecoder(file)
	if decorder != nil {
		var post2 Post
		err := decorder.Decode(&post2)
		if err == nil {
			fmt.Println("get post from file use json.NewDecoder", post2)
		}
	}
}

func main() {
	write2JsonFile()
	readFromJsonFile()
}
