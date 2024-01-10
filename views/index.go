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

func (*HtmlApi) Index(w http.ResponseWriter, r *http.Request) {
	index := commom.Template.Index
	// 涉及到上方所有数据结构的定义，都需要存在相关定义
	// 数据库查询的数据返回记录
	// 分页获取
	err2 := r.ParseForm()
	if err2 != nil {
		log.Println("表单获取失败", err2)
		index.WriteError(w, errors.New("The System is ERROR, Please contact the supervisor."))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Get All Index Information Error: ", err)
		index.WriteError(w, errors.New("The System is ERROR, Please contact the supervisor."))
	}
	index.WriteData(w, hr)
}
