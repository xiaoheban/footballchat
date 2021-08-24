package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"author"` //映射成数据库里面的字段 结构中的大写都会映射为数据库中的小写 db:"author"不能有空格
	//author string
}

var xDb *sqlx.DB

func init() {
	var err error
	xDb, err = sqlx.Open("postgres", "postgres://postgres:1234@47.103.40.211:8100/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (post *Post) create() (err error) {
	err = xDb.QueryRow("insert into posts (content,author) values ($1,$2) returning id",
		post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func getPost(id int) (post Post, err error) {
	err = xDb.QueryRowx("select id,content,author from posts where id=$1", id).StructScan(&post)
	return
}

/**
获取多篇帖子
*/
func getPosts() (posts []Post, err error) {
	rows, _ := xDb.Queryx("select * from posts")
	for rows.Next() {
		post := Post{}
		rows.StructScan(&post)
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func main() {
	post := Post{
		Content:    "今宵酒醒何处，杨柳岸晓风残月",
		AuthorName: "柳永",
	}
	err := post.create()
	if err != nil {
		fmt.Println(err)
		return
	}
	readPost, err := getPost(post.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("new ", readPost)
	allPosts, err := getPosts()
	if err != nil {
		fmt.Println("get posts failed")
		return
	}
	fmt.Println(allPosts)

}
