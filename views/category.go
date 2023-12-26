package views

import (
	"MyBlog/commom"
	"MyBlog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := commom.Template.Category
	path := r.URL.Path
	// println(path) // /c/1
	cIdString := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdString)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("此路径无法识别"))
		return
	}
	// 分区获取
	// 分页获取
	err2 := r.ParseForm()
	if err2 != nil {
		log.Println("表单获取失败", err2)
		categoryTemplate.WriteError(w, errors.New("The System is ERROR, Please contact the supervisor."))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10
	categoryResponse, _ := service.GetPostsByCategoryId(cId, page, pageSize)
	categoryTemplate.WriteData(w, categoryResponse)
}
