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

//-----------------/douyin/user/register/接口测试--------------------

// 用用户名为test1,密码为123456进行注册
func TestRegister(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/user/register/?username=test1&password=123456"
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
	assert.Equal(t, strings.Contains(str, "注册成功"), true) //如果是第一次注册则为true
}

// 重复用户名进行注册
func TestDisplayRegister_DuplicatedName(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/user/register/?username=test1&password=123"
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
	assert.Equal(t, strings.Contains(str, "用户名已存在"), true) //如果该用户名已经注册了则为true
}

// 用户名超过32位
func TestOver32Name(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/user/register/?username=123456789999999999999999999999999996666666&password=123"
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
	assert.Equal(t, strings.Contains(str, "用户名或密码不能超过32位"), true) //如果该用户名超过32位则为true
}

// 密码超过32位
func TestOver32Password(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/user/register/?username=test2&password=12345678999999999999999999999999999"
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
	assert.Equal(t, strings.Contains(str, "用户名或密码不能超过32位"), true) //如果密码超过32位则为true
}
