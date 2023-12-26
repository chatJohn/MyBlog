package dao

import "errors"

func GetUserNameById(uId int) (string, error) {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", uId)
	if row.Err() != nil {
		return "", errors.New("Query User Name Error")
	}
	var UserName string
	_ = row.Scan(&UserName)
	return UserName, nil
}
