package service

import (
	"MyBlog/config"
	"MyBlog/dao"
	"MyBlog/models"
	"html/template"
)

func GetPostDetail(pId int) (models.PostRes, error) {
	post, err := dao.GetPostDetailById(pId)
	var postRes models.PostRes
	if err != nil {
		return postRes, nil
	}
	cName, _ := dao.GetCategoryNameById(post.CategoryId)
	uName, _ := dao.GetUserNameById(post.UserId)
	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: cName,
		UserId:       post.UserId,
		UserName:     uName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}

	postRes.Viewer = config.Cfg.Viewer
	postRes.SystemConfig = config.Cfg.System
	postRes.Article = postMore
	return postRes, nil
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func SearchPost(searchValue string) []models.SearchResp {
	posts, _ := dao.GetPostSearch(searchValue)
	var searchRes []models.SearchResp
	for _, post := range posts {
		searchRes = append(searchRes, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	return searchRes
}
