package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func TestFileAction() {
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
	file1, _ := os.Create("osFile.txt") //如果文件已经存在会清空，否则会新建文件
	defer file1.Close()                 //用完需要关闭文件
	writtenLen, err := file1.Write(bytes)
	if err != nil {
		//发生了错误
		panic(err)
	}
	//读
	//接收缓冲区
	readBuffer := make([]byte, writtenLen)
	file2, _ := os.Open("osFile.txt")
	defer file2.Close()
	readLen, err := file2.Read(readBuffer)
	if err != nil {
		panic(err)
	}
	if readLen > 0 {
		fmt.Println("read file ", string(readBuffer))
	}
}
