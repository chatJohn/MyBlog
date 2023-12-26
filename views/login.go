package views

import (
	"MyBlog/commom"
	"MyBlog/config"
	"net/http"
)

func (*HtmlApi) Login(w http.ResponseWriter, r *http.Request) {
	login := commom.Template.Login
	login.WriteData(w, config.Cfg.Viewer)
}
