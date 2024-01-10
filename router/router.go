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
	// 发布或者更新文章
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	// 发布完成后返回文章详情到markdown格式
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)

	// 归档
	http.HandleFunc("/pigeonhole", views.HTML.PigeOnHole)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
