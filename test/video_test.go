package test

import (

	"ByteRhythm/app/video/service"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"net/http"
	"strings"
	"mime/multipart"
	"testing"
	"github.com/stretchr/testify/assert"
)

//-----------------/douyin/publish/action接口测试--------------------

//发布成功
func TestPublish(t *testing.T){
	url := "http://192.168.30.128:4000/douyin/publish/action"
	method := "POST"
	filePath := "/home/jan/Downloads/1.png" //这里的视频一定要确保自己电脑上有
	client := &http.Client{}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file,err := os.Open(filePath)
	assert.Empty(t,err)
	defer func(file *os.File){
		err := file.Close()
		assert.Empty(t,err)
	}(file)
	fileWriter, err := writer.CreateFormFile("data", file.Name())
	assert.Empty(t, err)

	_, err = io.Copy(fileWriter, file)
	assert.Empty(t, err)

	_ = writer.WriteField("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTI5VDIzOjE4OjA5LTA3OjAwIiwiZXhwIjoxNjkzMzgwMjEwLCJpYXQiOjE2OTMzNzY2MTAsImlzcyI6InRlc3QxIn0.Nx4DhtFFNzkP30NILuBlpKa_Qo2lb5MAkGp_uccij8I")
	_ = writer.WriteField("title", "只狼")

	err = writer.Close()
	assert.Empty(t, err)

	req, err := http.NewRequest(method, url, payload)
	assert.Empty(t, err)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	assert.Empty(t, err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Empty(t, err)
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	str := string(body)
	assert.Empty(t, err)
	video := service.GetVideoSrv()
	err = json.Unmarshal(body, &video)
	assert.Empty(t, err)
	assert.Equal(t, strings.Contains(str,"发布成功"),true)

}

//-----------------/douyin/feed接口测试--------------------

//获取视频流成功
func TestFeed(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/feed/?latest_time=2023-08-29T21:09:40.999Z&" +
		"token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTI5VDIzOjE4OjA5LTA3OjAwIiwiZXhwIjoxNjkzMzgwMjEwLCJpYXQiOjE2OTMzNzY2MTAsImlzcyI6InRlc3QxIn0.Nx4DhtFFNzkP30NILuBlpKa_Qo2lb5MAkGp_uccij8I"

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
	feed := service.GetVideoSrv()
	err = json.Unmarshal(body, &feed)
	assert.Empty(t, err)
	assert.Equal(t, strings.Contains(str,"获取视频流成功"),true)
}

//-----------------/douyin/publish/list接口测试--------------------

func TestPublishList(t *testing.T) {
	url := "http://192.168.30.128:4000/douyin/publish/list?user_id=1" +
		"token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTI5VDIzOjE4OjA5LTA3OjAwIiwiZXhwIjoxNjkzMzgwMjEwLCJpYXQiOjE2OTMzNzY2MTAsImlzcyI6InRlc3QxIn0.Nx4DhtFFNzkP30NILuBlpKa_Qo2lb5MAkGp_uccij8I"

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
	feed := service.GetVideoSrv()
	err = json.Unmarshal(body, &feed)
	assert.Empty(t, err)
	assert.Equal(t, strings.Contains(str, "获取成功"), true)
}