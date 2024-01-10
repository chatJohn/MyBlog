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

func GetPostCountBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(*) from blog_post where slug = ?", slug)
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
func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post")
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
func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug = ? limit ?, ?", slug, page, pageSize)
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

func GetPostDetailById(pId int) (models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ?", pId)
	var post models.Post
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(
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
		return post, err
	}
	return post, nil
}

func SavePost(post *models.Post) {
	result, err := DB.Exec("insert into blog_post"+
		"(title, content, markdown, category_id, user_id, view_count, type, slug, create_at, update_at) "+
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type,
		post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
	}

	pId, _ := result.LastInsertId()
	post.Pid = int(pId)
}
func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println(err)
	}
}
