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

// -----------------/douyin/comment/action/接口测试--------------------
// 发布评论
func TestCommentAction(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/comment/action/?"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OSwidXNlcm5hbWUiOiJ0ZXN0MiIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDAyOjIyOjMzLTA3OjAwIiwiZXhwIjoxNjkzMzg4MDAxLCJpYXQiOjE2OTMzODc5NzEsImlzcyI6InRlc3QyIn0.11RhlMGupHJDoetGndAWSiuNyBeAUHNEIF81VcVJkR0"
	video_id := "1"
	action_type := "1" //发布评论
	comment_text := "只狼太好玩了！"
	url := baseUrl + "token=" + token + "&video_id=" + video_id + "&action_type=" + action_type + "&comment_text=" + comment_text

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
	assert.Equal(t, strings.Contains(str, "评论成功"), true)

}

// 删除评论
func TestDeleteCommentAction(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/comment/action/?"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OSwidXNlcm5hbWUiOiJ0ZXN0MiIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDAyOjIyOjMzLTA3OjAwIiwiZXhwIjoxNjkzMzg4MDAxLCJpYXQiOjE2OTMzODc5NzEsImlzcyI6InRlc3QyIn0.11RhlMGupHJDoetGndAWSiuNyBeAUHNEIF81VcVJkR0"
	video_id := "1"
	action_type := "2" //删除评论
	comment_id := "1"
	url := baseUrl + "token=" + token + "&video_id=" + video_id + "&action_type=" + action_type + "&comment_id=" + comment_id

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
	assert.Equal(t, strings.Contains(str, "评论删除成功"), true)

}

//-----------------/douyin/comment/list接口测试--------------------

// 获取评论列表成功
func TestCommentList(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/comment/list/?"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OSwidXNlcm5hbWUiOiJ0ZXN0MiIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDAyOjIyOjMzLTA3OjAwIiwiZXhwIjoxNjkzMzg4MDAxLCJpYXQiOjE2OTMzODc5NzEsImlzcyI6InRlc3QyIn0.11RhlMGupHJDoetGndAWSiuNyBeAUHNEIF81VcVJkR0"
	video_id := "1"
	url := baseUrl + "token=" + token + "&video_id=" + video_id

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
	assert.Equal(t, strings.Contains(str, "获取评论列表成功"), true)

}
