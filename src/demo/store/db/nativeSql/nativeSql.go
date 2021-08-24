/**
原生sql的增删改查
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Post struct {
	id      int
	content string
	author  string
}

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:1234@47.103.40.211:8100/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

/**
创建一篇帖子
*/
func (p *Post) createPost() (err error) {
	sql := "insert into posts (content,author) values($1,$2) returning id" //预处理sql语句(模板)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.content, p.author).Scan(&p.id)

	return
}

/**
获取一篇单独的帖子
*/
func getPost(id int) (post Post, err error) {
	post = Post{}
	err = db.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.id, &post.content, &post.author)
	return
}

/*
更新帖子
*/
func (p *Post) update() error {
	_, err := db.Exec("update posts set author=$1,content=$2 where id=$3", p.author, p.content, p.id)
	return err
}

/**
删除帖子
*/
func (p *Post) delete() (err error) {
	_, err = db.Exec("delete from posts where id=$1", p.id)
	return
}

/**
获取很多帖子
*/
func getPosts(limit int) (posts []Post, err error) {
	rows, err := db.Query("select id,content,author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.id, &post.content, &post.author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}
func main() {

	post, err := getPost(4)
	if err != nil {
		return
	}
	fmt.Println(post)
	//post.delete()

	post.author = "浪子"
	post.update()

	//create post
	post1 := Post{
		author:  "西门吹风",
		content: "请问苏北往事什么时候更新啊？",
	}
	post1.createPost()
	fmt.Println("post1 ", post1)

	allPosts, err := getPosts(10)
	fmt.Println("10条数据 ", allPosts)
}
