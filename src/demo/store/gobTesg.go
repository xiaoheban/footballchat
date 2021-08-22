package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post1 struct {
	Id   int
	Name string
}

/**
通过gob写入
*/
func writeWithGob0() {
	post := Post1{
		Id:   1,
		Name: "jsim",
	}
	writeWithGob(post)
}
func writeWithGob(data interface{}) {
	fmt.Println(data)
	buff := new(bytes.Buffer) //等价于 buffer := & buffer.Buffer{}
	encoder := gob.NewEncoder(buff)
	err := encoder.Encode(data)
	if err == nil {
		err := ioutil.WriteFile("post", buff.Bytes(), 0600)
		if err == nil {
			fmt.Println("写入成功", err)
		} else {
			fmt.Println("写入失败")
		}
	}
}

/**
通过gob读取
*/
func readWithGob() {
	reads, _ := ioutil.ReadFile("post")
	buffer := bytes.NewBuffer(reads)
	decoder := gob.NewDecoder(buffer)
	var post Post1
	err := decoder.Decode(&post)
	if err == nil {
		fmt.Println(post)
	} else {
		fmt.Println("读取对象失败")
	}
}
