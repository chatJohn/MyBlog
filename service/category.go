package service

import (
	"MyBlog/config"
	"MyBlog/dao"
	"MyBlog/models"
	"errors"
	"html/template"
	"log"
)

func GetPostsByCategoryId(cId, page, pageSize int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println("Get All Category Error: ", err)
		return nil, err
	}
	posts, err2 := dao.GetPostPageByCategoryId(cId, page, pageSize)
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
	totalPostCnt := dao.GetPostCountByCategoryId(cId)
	var PageCount int = (totalPostCnt + pageSize - 1) / pageSize
	var pages []int
	for i := 1; i <= PageCount; i++ {
		pages = append(pages, i)
	}
	var hr *models.HomeResponse = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     totalPostCnt,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != PageCount,
	}
	categoryName, _ := dao.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		HomeResponse: hr,
		CategoryName: categoryName,
	}
	return categoryResponse, nil
}
