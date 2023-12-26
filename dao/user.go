package dao

import (
	"MyBlog/models"
	"errors"
	"log"
)

func GetUserNameById(uId int) (string, error) {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", uId)
	if row.Err() != nil {
		return "", errors.New("Query User Name Error")
	}
	var UserName string
	_ = row.Scan(&UserName)
	return UserName, nil
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ? limit 1", userName, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{} // *model.User ====> 这样会先初始化一个空指针
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
