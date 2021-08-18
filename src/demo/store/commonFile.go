package main

import (
	"fmt"
	"io/ioutil"
)

func testFileAction() {
	//准备数据
	bytes := []byte("hello world")
	//ioutil读写
	//写
	err := ioutil.WriteFile("iofile.txt", bytes, 0644)
	if err != nil {
		panic(err)
	}
	//读
	reads, err := ioutil.ReadFile("iofile.txt")
	if err != nil {
		fmt.Println(string(reads))
	}
	//os读写
	//写

	//读
}
