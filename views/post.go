package views

import (
	"MyBlog/commom"
	"MyBlog/service"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (*HtmlApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := commom.Template.Detail
	// 获取路径参数
	path := r.URL.Path
	trimPath := strings.TrimPrefix(path, "/p/")     // ==>7.html
	pIdStr := strings.TrimSuffix(trimPath, ".html") // ===> 7
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("路径参数出现错误，无法解析"))
		return
	}
	postDetail, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询错误"))
		return
	}
	detail.WriteData(w, postDetail)
}
