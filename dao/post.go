package dao

import (
	"MyBlog/models"
	"log"
)

func GetPostCount() (count int) {
	rows := DB.QueryRow("select count(*) from blog_post")
	rows.Scan(&count)
	return
}
func GetPostCountByCategoryId(cid int) (count int) {
	rows := DB.QueryRow("select count(*) from blog_post where category_id = ?", cid)
	rows.Scan(&count)
	return
}
func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
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

func GetPostPageByCategoryId(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?, ?", cid, page, pageSize)
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
