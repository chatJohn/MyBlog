package service

import (
	"MyBlog/dao"
	"MyBlog/models"
	"errors"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号或密码不正确")
	}
	loginRes := &models.LoginRes{}
	return loginRes, nil
}
