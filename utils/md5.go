package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Md5Crypt 给字符串生成md5
// @params str 需要加密的字符串
// @params salt interface{} 加密的盐
// @return str 返回md5码

func Md5Crypt(str string, salt ...interface{}) (CryptStr string) { // 在获取到前端传过来的密码的时候，密码已经被MD5加密，这里在进行MD5加密，然后到数据库进行查询
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
