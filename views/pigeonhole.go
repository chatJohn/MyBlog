package views

import (
	"MyBlog/commom"
	"MyBlog/service"
	"net/http"
)

func (*HtmlApi) PigeOnHole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := commom.Template.Pigeonhole
	pigeonholeRes := service.FindPostPigeOnHole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
