package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	id       int
	content  string
	author   string
	comments []Comment //切片保存的是引用 不是副本 //一个帖子有多个评论
}
type Comment struct {
	id      int
	content string
	author  string
	post    *Post //属于哪一个帖子
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:1234@47.103.40.211:8100/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

/**
创建帖子
*/
func (p *Post) create() (err error) {
	err = db.QueryRow("insert into posts (content,author) values ($1,$2) "+
		"returning id", p.content, p.author).Scan(&p.id)
	return
}

/**
创建一条评论
*/
func (c *Comment) create() (err error) {
	if c.post == nil {
		return
	}
	err = db.QueryRow("insert into comments (content,author,post_id) values ($1,$2,$3) returning id",
		c.content, c.content, c.post.id).Scan(&c.id)
	return
}

func getPostById(id int) (p Post, err error) {
	err = db.QueryRow("select id,author,content from posts where id=$1", id).Scan(&p.id, &p.author, &p.content)
	if err != nil {
		return
	}
	rows, err := db.Query("select id,author,content from comments where post_id=$1", p.id)
	if err != nil {
		return
	}
	//所有评论都赋给对应的帖子
	for rows.Next() {
		comment := Comment{
			post: &p,
		}
		err = rows.Scan(&comment.id, &comment.author, &comment.content)
		if err != nil {
			return
		}
		p.comments = append(p.comments, comment)
	}
	rows.Close() //记得要关闭
	return
}
func addComment2Post(id int) {
	post, err := getPostById(id)
	if err != nil {
		return
	}
	//发表评论
	comment := Comment{
		content: "太白威武，好诗！！",
		author:  "杜甫",
		post:    &post,
	}
	err = comment.create()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	/*	post := Post{
			author:  "李白",
			content: "日照香炉生紫烟，遥看瀑布挂前川",
		}
		fmt.Println("帖子初始值", post)
		err := post.create()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("帖子", post)

		//发表评论
		comment := Comment{
			content: "太白威武，好诗！！",
			author:  "杜甫",
			post:    &post,
		}
		err = comment.create()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("评论", comment)*/
	addComment2Post(1)
	//获取帖子
	readPost, err := getPostById(1)
	if err == nil {
		fmt.Println("帖子+评论", readPost)
	}
}
