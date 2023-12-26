package api

import (
	"MyBlog/commom"
	"MyBlog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	params := commom.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		commom.Error(w, err)
		return
	}
	commom.Success(w, loginRes)
}
