package api

import (
	"MyBlog/commom"
	"MyBlog/dao"
	"MyBlog/models"
	"MyBlog/service"
	"MyBlog/utils"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		commom.Error(w, errors.New("登录过期"))
		return
	}
	uid := claims.Uid
	// post方法是保存文章，put表示更新文章
	method := r.Method
	switch method {
	case http.MethodPost:
		{
			params := commom.GetRequestJsonParam(r)
			cId := params["categoryId"].(string)
			categoryId, _ := strconv.Atoi(cId)
			content := params["content"].(string)
			markdown := params["markdown"].(string)
			slug := params["slug"].(string)
			title := params["title"].(string)
			postType := params["type"].(float64)
			pType := int(postType)

			post := &models.Post{
				Pid:        0,
				Title:      title,
				Slug:       slug,
				Content:    content,
				Markdown:   markdown,
				CategoryId: categoryId,
				UserId:     uid,
				ViewCount:  0,
				Type:       pType,
				CreateAt:   time.Now(),
				UpdateAt:   time.Now(),
			}
			service.SavePost(post)
			commom.Success(w, post)
		}

	case http.MethodPut:
		params := commom.GetRequestJsonParam(r)
		cId := params["categoryId"].(float64)
		categoryId := int(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		commom.Success(w, post)
	}
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		commom.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostDetailById(pId)
	if err != nil {
		commom.Error(w, err)
		return
	}
	commom.Success(w, post)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	searchValue := r.Form.Get("val")
	searchRes := service.SearchPost(searchValue)
	commom.Success(w, searchRes)
}
