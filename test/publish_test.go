package test

import (
	"ByteRhythm/app/video/service"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"
)

//-----------------/douyin/publish/action/接口测试--------------------

// 发布成功
func TestPublish(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/publish/action/"
	method := "POST"
	filePath := "/home/jan/Downloads/1.png" //这里的视频一定要确保自己电脑上有
	client := &http.Client{}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(filePath)
	assert.Empty(t, err)
	defer func(file *os.File) {
		err := file.Close()
		assert.Empty(t, err)
	}(file)
	fileWriter, err := writer.CreateFormFile("data", file.Name())
	fmt.Printf(file.Name())
	assert.Empty(t, err)

	_, err = io.Copy(fileWriter, file)
	assert.Empty(t, err)

	_ = writer.WriteField("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJuYW1lIjoidGVzdDEiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0zMFQwNjowMzowMy0wNzowMCIsImV4cCI6MTY5MzQwNDY2OCwiaWF0IjoxNjkzNDA0NTQ4LCJpc3MiOiJ0ZXN0MSJ9.2dlgE1GpvyIjvXTK5Jl9d4ayQm2Rd9Czhf1tGMzue4A")
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
	//fmt.Printf(str)
	assert.Empty(t, err)
	video := service.GetVideoSrv()
	err = json.Unmarshal(body, &video)
	assert.Empty(t, err)
	assert.Equal(t, strings.Contains(str, "发布成功"), true)

}

//-----------------/douyin/publish/list/接口测试--------------------

// 获取成功
func TestPublishList(t *testing.T) {
	url := "http://192.168.256.128:8080/douyin/publish/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDA3OjI5OjAwLTA3OjAwIiwiZXhwIjoxNjkzNDA3ODEwLCJpYXQiOjE2OTM0MDc2OTAsImlzcyI6InRlc3QxIn0.hFbVJbjg-Ec9sdR0P1ufO4SM4goIU4njlPHMVScHZ0s&user_id=1"
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
