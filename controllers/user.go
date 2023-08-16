package controllers

import (
	"ByteRhythm/models"
	"ByteRhythm/utils"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	baseController
}

// Register 用户注册
func (c *UserController) Register() {

	username := c.GetString("username")
	password := c.GetString("password")

	if len(username) > 32 || len(password) > 32 {
		c.MsgFail("用户名或密码不能超过32位")
		return
	}
	//数据库校验用户名是否唯一
	if exist := c.o.QueryTable(new(models.User)).Filter("username", username).Exist(); exist {
		c.MsgFail("用户名已存在")
		return
	}

	avatar, _ := web.AppConfig.String("avatar")
	background, _ := web.AppConfig.String("background")
	user := models.User{
		Username:        username,
		Password:        utils.Md5(password),
		Avatar:          avatar,
		BackgroundImage: background,
	}

	if id, err := c.o.Insert(&user); err != nil {
		c.MsgFail("注册失败")
		return
	} else {
		token := utils.GenerateToken(user, 0)
		c.Data["json"] = map[string]interface{}{
			"status_code": 0,
			"status_msg":  "注册成功",
			"user_id":     id,
			"token":       token,
		}
		c.ServeJSON()
		return
	}
}

// MsgFail 返回用户注册失败信息
func (c *UserController) MsgFail(msg string) {
	c.Data["json"] = map[string]interface{}{
		"status_code": 1,
		"status_msg":  msg,
		"user_id":     nil,
		"token":       nil,
	}
	c.ServeJSON()
}

// Login 用户登录
func (c *UserController) Login() {
	var user models.User
	username := c.GetString("username")
	password := utils.Md5(c.GetString("password"))
	if len(username) > 32 || len(password) > 32 {
		c.MsgFail("用户名或密码不能超过32位")
		return
	}
	query := c.o.QueryTable(new(models.User)).Filter("username", username)
	if exist := query.Exist(); !exist {
		c.MsgFail("用户未注册")
		return
	}
	if err := query.Filter("password", password).One(&user); err != nil {
		c.MsgFail("用户名或密码错误")
		return
	} else {
		token := utils.GenerateToken(user, 0)
		c.Data["json"] = map[string]interface{}{
			"status_code": 0,
			"status_msg":  "登录成功",
			"user_id":     user.Id,
			"token":       token,
		}
		c.ServeJSON()
		return
	}
}

// Info 获取用户信息
func (c *UserController) Info() {
	uid, _ := c.GetInt("user_id")
	token := c.GetString("token")

	userInfo := c.GetUserInfo(uid, token)
	c.Data["json"] = map[string]interface{}{
		"status_code": 0,
		"status_msg":  "获取用户信息成功",
		"user":        userInfo,
	}
	c.ServeJSON()
	return
}
