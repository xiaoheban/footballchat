package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//帖子
type Post struct {
	Id       int
	Author   string `sql:"not null"`
	Content  string
	Coments  []Comment //评论列表
	CreateAt time.Time //创建时间，设置成这个会自动创建这个时间
}

//评论
type Comment struct {
	Id      int
	Author  string `sql:"not null"`
	PostId  int    `sql:"index"` //属性名是固定的吗 index起到什么作用
	Content string
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp ssmode=disable")
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Post{}, &Comment{}) //可以自动迁移数据库，如果Post或Comment增减字段，执行这个以后会自动更新表字段
}

func main() {
	//增
	post := Post{
		Author:  "jim",
		Content: "goOrm创建的第一个帖子",
	}
	DB.Create(&post)
	//创建评论
	comment := Comment{
		Author:  "jack",
		Content: "goorm评论",
	}
	//关联评论和帖子
	DB.Model(&post).Association("Comments").Append(comment)
	//删
	DB.Delete()
	//改
	//查
	DB.Where()
}
