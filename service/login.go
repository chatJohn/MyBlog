package service

import (
	"MyBlog/dao"
	"MyBlog/models"
	"MyBlog/utils"
	"errors"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	println("passwd = ", passwd)
	passwd = utils.Md5Crypt(passwd, "chatting")
	println("passwd = ", passwd)
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号或密码不正确")
	}
	uID := user.Uid
	// 生成token,jwt技术进行生成令牌
	award, err := utils.Award(&uID)
	if err != nil {
		return nil, errors.New("token未能生成")
	}

	var userInfo models.UserInfo
	userInfo.Uid = uID
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	loginRes := &models.LoginRes{
		Token:    award,
		UserInfo: userInfo,
	}
	return loginRes, nil
}
