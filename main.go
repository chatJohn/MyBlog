package main

import (
	"MyBlog/commom"
	"MyBlog/router"
	"log"
	"net/http"
)

type IndexData struct {
	Title       string `json:"title"` // 返回json格式的数据的时候，显示相关字段名
	Description string `json:"description"`
}

func init() {
	// 模板加载
	commom.LoadTemplate()
}
func main() {
	// 一个项目只能有一个main
	server := http.Server{Addr: "127.0.0.1:8080"}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
