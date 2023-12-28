package views

import (
	"MyBlog/commom"
	"MyBlog/service"
	"net/http"
)

func (*HtmlApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := commom.Template.Writing
	// write data
	write := service.Writing()
	writing.WriteData(w, write)
}
