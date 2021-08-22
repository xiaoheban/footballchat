package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	id   int
	name string
}

const (
	CSV_FILE_NAME = "test.csv"
)

/**
读写csv文件
*/
func writeAndReadCsv() bool {
	//待写入的数据
	posts := []Post{
		{
			id:   1,
			name: "jim",
		},
		{2, "jack"},
		{3, "tom"},
	}
	file, err := os.Create(CSV_FILE_NAME)
	if err != nil {
		return false
	}
	defer file.Close()
	writer := csv.NewWriter(file) //write实现了write接口
	for _, post := range posts {
		record := []string{strconv.Itoa(post.id), post.name}
		err := writer.Write(record)
		if err != nil {
			return false
		}
	}
	writer.Flush() //防止有数据在缓冲区中没有写入文件，
	//可以调用此方法，一般也可以不调用，一般是打开一个文件进行写操作，接着对文件进行读操作需要这样
	return true
}

/**
从csv读取数据
*/
func readFromCsv() {
	file, err := os.Open(CSV_FILE_NAME)
	if err != nil {
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	//设置成-1 即使每次读到的记录缺少某些字段也不会报错，
	//如果设置成正数，则每次读取的数据有空的属性，则会报错，
	//如果设置成0，则设置成第一次读取记录的字段数量
	records, err := reader.ReadAll()
	//对于数据不多的可以直接用readall,如果数据比较多则用Read多次读取
	if err != nil {
		return
	}
	var posts []Post
	for _, record := range records {
		id, _ := strconv.ParseInt(record[0], 0, 0)
		post := Post{
			id:   int(id),
			name: record[1],
		}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].name)
}
