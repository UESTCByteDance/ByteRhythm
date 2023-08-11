package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//func GenerateToken(id int) string {
//	// 设置 JWT 密钥
//	secretKey := "hwifauejanwjjhafaoiweflajf"
//
//	// 创建一个新的 Token
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	// 设置 Token 的有效期
//	expirationTime := time.Now().Add(30 * time.Minute) // 有效期为 30分钟
//	claims := token.Claims.(jwt.MapClaims)
//	claims["exp"] = expirationTime.Unix()
//
//	//在这里可以设置其他自定义的声明，例如用户 ID 等
//	claims["id"] = id
//
//	// 使用密钥签署 Token，并获取字符串形式的 Token
//	tokenString, err := token.SignedString([]byte(secretKey))
//	if err != nil {
//		panic(err)
//	}
//	return tokenString
//}
//
//func LoginFilter(ctx *context.Context) {
//	token := ctx.Input.Session("token")
//	if token == nil {
//		//返回json数据
//		ctx.Output.JSON(map[string]interface{}{
//			"status_code": 1,
//			"status_msg":  "未登录",
//			"user":        nil,
//		}, false, false)
//	}
//}
