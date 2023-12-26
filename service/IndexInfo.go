package service

import (
	"MyBlog/config"
	"MyBlog/dao"
	"MyBlog/models"
	"errors"
	"html/template"
	"log"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("Get All Category Error: ", err)
		return nil, err
	}
	posts, err2 := dao.GetPostPage(page, pageSize)
	if err2 != nil {
		log.Println("Get All Post Error: ", err2)
		return nil, err2
	}
	var postMores []models.PostMore
	for _, post := range posts {
		cName, err3 := dao.GetCategoryNameById(post.CategoryId)
		uName, err4 := dao.GetUserNameById(post.UserId)
		if err3 != nil || err4 != nil {
			return nil, errors.New(err3.Error() + err4.Error())
		}
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: cName,
			UserId:       post.UserId,
			UserName:     uName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	var hr *models.HomeResponse = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	return hr, nil
}
