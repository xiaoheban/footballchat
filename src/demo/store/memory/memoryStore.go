package memory

import "fmt"

/**
帖子数据结构
*/
type Post struct {
	ID      int
	Author  string
	Content string
}

var idPost map[int]*Post          //存储id和帖子
var authorPost map[string][]*Post //作者帖子映射，每个作者有多个帖子
/**
把数据存储到容器里面
*/
func store(post Post) {
	idPost[post.ID] = &post
	authorPost[post.Author] = append(authorPost[post.Author], &post)
}
func Test() {
	post1 := Post{
		ID:      1,
		Author:  "jim",
		Content: "I am jim",
	}
	post2 := Post{
		ID:      2,
		Author:  "jim1",
		Content: "I am jim1",
	}
	post3 := Post{
		ID:      3,
		Author:  "jim2",
		Content: "I am jim2",
	}
	post4 := Post{
		ID:      4,
		Author:  "jim3",
		Content: "I am jim3",
	}
	idPost = make(map[int]*Post)
	authorPost = make(map[string][]*Post)
	store(post1)
	store(post2)
	store(post3)
	store(post4)
	//根据id查找帖子
	fmt.Println(idPost[1])
	fmt.Println(idPost[2])
	//根据作者查找帖子，是一个数组
	posts := authorPost["jim"]
	//遍历帖子数组
	for _, post := range posts {
		fmt.Println(post.Content)
	}
}
