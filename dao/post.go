package dao

import (
	"MyBlog/models"
	"log"
)

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post limit ?, ?", page, pageSize)
	if err != nil {
		log.Println("Query Post Error: ", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post = models.Post{}
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil

}
