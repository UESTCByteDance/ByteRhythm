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

//--------/douyin/user/接口测试--------------------------

// 获取用户信息成功，是test1用户的
func TestUserInfo(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/user/?user_id=10&token=" +
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAsInVzZXJuYW1lIjoidGVzdDEiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0zMFQwNDo0Njo1OC0wNzowMCIsImV4cCI6MTY5MzM5NjQzOCwiaWF0IjoxNjkzMzk2MzE4LCJpc3MiOiJ0ZXN0MSJ9.NQymZNYRryaBFpbaYApSqBuwKfyYEIIGHLZGRNEn1as"
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
	assert.Equal(t, strings.Contains(str, "获取用户信息成功"), true)
}

// 用户不存在
func TestNoUser(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/user/?user_id=14&token=" +
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAsInVzZXJuYW1lIjoidGVzdDEiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0zMFQwNDo0Njo1OC0wNzowMCIsImV4cCI6MTY5MzM5NjQzOCwiaWF0IjoxNjkzMzk2MzE4LCJpc3MiOiJ0ZXN0MSJ9.NQymZNYRryaBFpbaYApSqBuwKfyYEIIGHLZGRNEn1as"
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
	assert.Equal(t, strings.Contains(str, "用户不存在"), true)
}
