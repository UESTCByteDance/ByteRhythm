package test

import (
	"ByteRhythm/app/user/service"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

//-----------------/douyin/user/login/接口测试--------------------

// 登陆成功
func TestLogin(t *testing.T) {

	url := "http://192.168.256.128:8080/douyin/user/login/?username=test1&password=123456"
	method := "POST"
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
	assert.Equal(t, strings.Contains(str, "登录成功"), true) //登录成功则为true
}

// 用户名为空进行登录
func TestLogin_EmptyName(t *testing.T) {

	url := "http://192.168.256.128:8080/douyin/user/login/?username=&password=123456"
	method := "POST"
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
	assert.Equal(t, strings.Contains(str, "用户名或密码错误"), true)
}

// 密码为空进行登录
func TestLogin_EmptyPassword(t *testing.T) {

	url := "http://192.168.256.128:8080/douyin/user/login/?username=test1&password="
	method := "POST"
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
	assert.Equal(t, strings.Contains(str, "用户名或密码错误"), true)
}

// 未注册的用户名登录
func TestLogin_NoRegister(t *testing.T) {

	url := "http://192.168.256.128:8080/douyin/user/login/?username=test3&password=123"
	method := "POST"
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
	assert.Equal(t, strings.Contains(str, "用户名或密码错误"), true)
}
