package service

import (
	"MyBlog/config"
	"MyBlog/dao"
	"MyBlog/models"
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
	var hr *models.HomeResponse = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	return hr, nil
}
