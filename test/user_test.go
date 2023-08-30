package test

import (

	"ByteRhythm/app/user/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)
//-----------------/douyin/user/register接口测试--------------------

//用用户名为test1,密码为123456进行注册
func TestRegister(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/user/register?username=test1&password=123456"
	method := "POST"

	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"注册成功"),true)//如果是第一次注册则为true
}

// 重复用户名进行注册
func TestDisplayRegister_DuplicatedName(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/user/register?username=test1&password=123"
	method := "POST"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"用户名已存在"),true)//如果该用户名已经注册了则为true
}

//用户名超过32位
func TestOver32Name(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/user/register?username=12345678999999999999999999999999999&password=123"
	method := "POST"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"用户名或密码不能超过32位"),true)//如果该用户名超过32位则为true
}

//密码超过32位
func TestOver32Password(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/user/register?username=test2&password=12345678999999999999999999999999999"
	method := "POST"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"用户名或密码不能超过32位"),true)//如果密码超过32位则为true
}

//-----------------/douyin/user/login接口测试--------------------

// 登陆成功
func TestLogin(t *testing.T) {

	url := "http://192.168.30.128:4000/douyin/user/login?username=test1&password=123456"
	method := "POST"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"登录成功"),true)//登录成功则为true
}

//用户名为空进行登录
func TestLogin_EmptyName(t *testing.T) {

	url := "http://192.168.30.128:4000/douyin/user/login?username=&password=123456"
	method := "POST"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"用户名或密码错误"),true)
}

//密码为空进行登录
func TestLogin_EmptyPassword(t *testing.T) {

	url := "http://192.168.30.128:4000/douyin/user/login?username=test1&password="
	method := "POST"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"用户名或密码错误"),true)
}

//未注册的用户名登录
func TestLogin_NoRegister(t *testing.T) {

	url := "http://192.168.30.128:4000/douyin/user/login?username=test3&password=123"
	method := "POST"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"用户名或密码错误"),true)
}

//--------/douyin/user接口测试--------------------------

//获取用户信息成功，是test1用户的
func TestUserInfo(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/user?user_id=12&token=" +
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIsInVzZXJuYW1lIjoidGVzdDEiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0yOVQyMDoxOTozOC0wNzowMCIsImV4cCI6MTY5MzM2OTc3NiwiaWF0IjoxNjkzMzY2MTc2LCJpc3MiOiJ0ZXN0MSJ9.pbZ7_IA8xlXCZXyAhnceyvZ5LFlZ63AWC3_TQwK5XIw"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"获取用户信息成功"),true)
}

//用户不存在
func TestNoUser(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/user?user_id=14&token=" +
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIsInVzZXJuYW1lIjoidGVzdDEiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0yOVQyMDoxOTozOC0wNzowMCIsImV4cCI6MTY5MzM2OTc3NiwiaWF0IjoxNjkzMzY2MTc2LCJpc3MiOiJ0ZXN0MSJ9.pbZ7_IA8xlXCZXyAhnceyvZ5LFlZ63AWC3_TQwK5XIw"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(body)
	fmt.Println(str)
	user := service.GetUserSrv()
	err = json.Unmarshal(body, &user)
	assert.Empty(t, err)
	assert.Equal(t,strings.Contains(str,"用户不存在"),true)
}