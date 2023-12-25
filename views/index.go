package views

import (
	"MyBlog/commom"
	"MyBlog/config"
	"MyBlog/models"
	"net/http"
)

func (*HtmlApi) Index(w http.ResponseWriter, r *http.Request) {
	index := commom.Template.Index
	// 涉及到上方所有数据结构的定义，都需要存在相关定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "chatting",
			ViewCount:    123,
			CreateAt:     "2023-12-24",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
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
	index.WriteData(w, hr)
}
