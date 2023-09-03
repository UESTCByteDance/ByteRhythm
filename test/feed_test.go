package test

import (
	"ByteRhythm/app/video/service"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

//-----------------/douyin/feed/接口测试--------------------

// 获取视频流成功
func TestFeed(t *testing.T) {

	url := "http://192.168.256.128:8080/douyin/feed/?latest_time=2023-08-30%2007%3A56%3A40.213&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0MSIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTMwVDA3OjI5OjAwLTA3OjAwIiwiZXhwIjoxNjkzNDA3NzEyLCJpYXQiOjE2OTM0MDc1OTIsImlzcyI6InRlc3QxIn0.ns00VSFt3dGUYY73h4P_njeP6F17HahTlmLJ7sJZWjM"
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
	str := string(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	assert.Empty(t, err)
	video := service.GetVideoSrv()
	err = json.Unmarshal(body, &video)
	assert.Empty(t, err)
	assert.Equal(t, strings.Contains(str, "获取视频流成功"), true)
}
