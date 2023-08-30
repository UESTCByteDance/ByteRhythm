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
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTI5VDIzOjE4OjA5LTA3OjAwIiwiZXhwIjoxNjkzMzgwMjEwLCJpYXQiOjE2OTMzNzY2MTAsImlzcyI6InRlc3QxIn0.Nx4DhtFFNzkP30NILuBlpKa_Qo2lb5MAkGp_uccij8I"
	//token随时都在变，要测的时候再具体修改
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
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTI5VDIzOjE4OjA5LTA3OjAwIiwiZXhwIjoxNjkzMzgwMjEwLCJpYXQiOjE2OTMzNzY2MTAsImlzcyI6InRlc3QxIn0.Nx4DhtFFNzkP30NILuBlpKa_Qo2lb5MAkGp_uccij8I"
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