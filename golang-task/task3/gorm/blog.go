package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，
// 其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
// Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	Id         int    `gorm:"primary_key;AUTO_INCREMENT"`
	Username   string `gorm:"type:varchar(100);not null"`
	Email      string `gorm:"type:varchar(100);not null"`
	Posts      []Post `gorm:"foreignKey:UserId"`
	PostsCount int    `gorm:"default:0"` // 可选：用于统计用户的文章数量
}

type Post struct {
	Id       int       `gorm:"primary_key;AUTO_INCREMENT"`
	Title    string    `gorm:"type:varchar(100);not null"`
	Content  string    `gorm:"type:text;not null"`
	UserId   int       `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:PostId"`
	//评论数量
	CommentsStatus string `gorm:"type:varchar(20);default:'有评论'"`
	//评论状态
}

type Comment struct {
	Id      int    `gorm:"primary_key;AUTO_INCREMENT"`
	Content string `gorm:"type:text;not null"`
	PostId  int    `gorm:"not null"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移模型到数据库
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	//构造一些初始话数据
	// user1 := User{
	// 	Email:    "test@example.com",
	// 	Username: "test",
	// }
	// db.Create(&user1) // 创建用户
	// post := Post{
	// 	Title:   "My Second Post",
	// 	Content: "This is the content of my second post.",
	// 	UserId:  1, // 假设用户ID为1
	// }
	// db.Create(&post) // 创建文章

	// comment := Comment{
	// 	Content: "This is a comment.",
	// 	PostId:  1, // 假设文章ID为1Second
	// }
	// db.Create(&comment) // 创建评论

	//关联查询
	// 基于上述博客系统的模型定义。
	// 要求 ：
	// // 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// var user User
	// db.Preload("Posts.Comments").First(&user, 1) // 假设用户ID为1
	// // 输出用户信息和文章及评论
	// for _, post := range user.Posts {
	// 	println("Post Title:", post.Title)
	// 	for _, comment := range post.Comments {
	// 		println("Comment:", comment.Content)
	// 	}
	// }
	// // 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	// var postWithMostComments Post
	// db.Raw(`SELECT p.* FROM posts p
	// JOIN comments c ON p.id = c.post_id
	// GROUP BY p.id
	// ORDER BY COUNT(c.id) DESC
	// LIMIT 1`).Scan(&postWithMostComments)
	// println("Post with most comments:", postWithMostComments.Title)

	//删除文章评论
	db.Debug().Delete(&Comment{Id: 4}) // 假设删除文章ID为1的所有评论
}

// 钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 假设 User 模型中有一个字段 PostsCount 用于统计文章数量
	if err := tx.Model(&User{}).Where("id=?", p.UserId).Update("posts_count", gorm.Expr("posts_count+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
// 如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	// 在删除前保存 PostId
	var comment Comment
	if err := tx.Where("id = ?", c.Id).First(&comment).Error; err != nil {
		return err
	}

	// 保存 PostId 以便在删除后使用
	postId := comment.PostId

	// 在事务结束后执行检查逻辑
	tx.InstanceSet("post_id", postId)

	return nil
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 获取之前保存的 PostId
	postIdValue, _ := tx.InstanceGet("post_id")
	var postId int

	if postIdValue != nil {
		postId = postIdValue.(int)
	} else {
		// 备用方案：如果无法获取 PostId，则不执行更新操作
		return nil
	}

	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", postId).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		// 假设 Post 模型中有一个字段 CommentsStatus 用于表示评论状态
		if err := tx.Model(&Post{}).Where("id = ?", postId).Update("comments_status", "无评论").Error; err != nil {
			return err
		}
	}
	return nil
}
