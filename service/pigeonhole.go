package service

import (
	"MyBlog/config"
	"MyBlog/dao"
	"MyBlog/models"
)

func FindPostPigeOnHole() models.PigeOnHoleRes {
	// 查询所有的文章 进行月份的整理
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	// 查询所有的分类
	category, _ := dao.GetAllCategory()

	return models.PigeOnHoleRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    category,
		Lines:        pigeonholeMap,
	}
}
