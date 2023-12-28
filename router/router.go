package router

import (
	"MyBlog/api"
	"MyBlog/views"
	"net/http"
)

func Router() {
	// 1. 页面路由 2. 数据路由(json) 3. 静态资源路由
	http.HandleFunc("/", views.HTML.Index)
	// 分类http://localhost:8080/c/1
	http.HandleFunc("/c/", views.HTML.Category)
	// 登录
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/api/v1/login", api.API.Login)
	// 文章详情
	http.HandleFunc("/p/", views.HTML.Detail)

	// 用户登录网站之后，可以写文章
	http.HandleFunc("/writing", views.HTML.Writing)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
