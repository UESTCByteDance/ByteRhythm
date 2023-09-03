package test

import (
	"ByteRhythm/app/favorite/service"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

// -----------------/douyin/message/chat/接口测试--------------------
// 获取聊天记录成功！
func TestChatMessage(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/message/chat/?"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OSwidXNlcm5hbWUiOiJ0ZXN0MiIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDAyOjIyOjMzLTA3OjAwIiwiZXhwIjoxNjkzMzg4MDAxLCJpYXQiOjE2OTMzODc5NzEsImlzcyI6InRlc3QyIn0.11RhlMGupHJDoetGndAWSiuNyBeAUHNEIF81VcVJkR0"
	to_user_id := "2"
	pre_msg_time := "2023-08-30%2023%3A32%3A49"
	url := baseUrl + "token=" + token + "&to_user_id=" + to_user_id + "&pre_msg_time=" + pre_msg_time

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	assert.Empty(t, err)

	res, err := client.Do(req)
	assert.Empty(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Empty(t, err)
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	str := string(body)
	assert.Empty(t, err)
	favor := service.GetFavoriteSrv()
	err = json.Unmarshal(body, &favor)
	assert.Empty(t, err)
	fmt.Printf(string(body))
	assert.Equal(t, strings.Contains(str, "获取聊天记录成功！"), true)

}

// 不能查看与自己的聊天记录！
func TestChatMessageMyself(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/message/chat/?"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OSwidXNlcm5hbWUiOiJ0ZXN0MiIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDAyOjIyOjMzLTA3OjAwIiwiZXhwIjoxNjkzMzg4MDAxLCJpYXQiOjE2OTMzODc5NzEsImlzcyI6InRlc3QyIn0.11RhlMGupHJDoetGndAWSiuNyBeAUHNEIF81VcVJkR0"
	to_user_id := "1"
	pre_msg_time := "2023-08-30%2023%3A32%3A49"
	url := baseUrl + "token=" + token + "&to_user_id=" + to_user_id + "&pre_msg_time=" + pre_msg_time

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	assert.Empty(t, err)

	res, err := client.Do(req)
	assert.Empty(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Empty(t, err)
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	str := string(body)
	assert.Empty(t, err)
	favor := service.GetFavoriteSrv()
	err = json.Unmarshal(body, &favor)
	assert.Empty(t, err)
	fmt.Printf(string(body))
	assert.Equal(t, strings.Contains(str, "不能查看与自己的聊天记录！"), true)

}

// -----------------/douyin/message/action/接口测试--------------------
// 发送消息成功！
func TestActionMessage(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/message/action/?"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OSwidXNlcm5hbWUiOiJ0ZXN0MiIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDAyOjIyOjMzLTA3OjAwIiwiZXhwIjoxNjkzMzg4MDAxLCJpYXQiOjE2OTMzODc5NzEsImlzcyI6InRlc3QyIn0.11RhlMGupHJDoetGndAWSiuNyBeAUHNEIF81VcVJkR0"
	to_user_id := "2"
	action_type := "1"
	content := "你弦一郎打了多少遍？"
	url := baseUrl + "token=" + token + "&to_user_id=" + to_user_id + "&action_type=" + action_type + "&content=" + content

	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	assert.Empty(t, err)

	res, err := client.Do(req)
	assert.Empty(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Empty(t, err)
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	str := string(body)
	assert.Empty(t, err)
	favor := service.GetFavoriteSrv()
	err = json.Unmarshal(body, &favor)
	assert.Empty(t, err)
	fmt.Printf(string(body))
	assert.Equal(t, strings.Contains(str, "发送消息成功！"), true)

}
