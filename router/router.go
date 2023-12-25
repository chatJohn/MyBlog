package router

import (
	"MyBlog/views"
	"net/http"
)

func Router() {
	// 1. 页面路由 2. 数据路由(json) 3. 静态资源路由
	http.HandleFunc("/", views.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
