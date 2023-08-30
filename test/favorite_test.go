package test

import (

	"ByteRhythm/app/favorite/service"
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

//-----------------/douyin/favorite/action接口测试--------------------
//点赞成功
func TestFavoriteAction(t *testing.T) {
	baseUrl :="http://192.168.30.128:4000/douyin/favorite/action?"
	video_id := "1"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ0ZXN0MyIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTI5VDIzOjI5OjA1LjY1Ni0wNzowMCIsImV4cCI6MTY5MzM4MDU0NSwiaWF0IjoxNjkzMzc2OTQ1LCJpc3MiOiJ0ZXN0MyJ9.hZxgHYARp_xs9hCylBav9YYwFqkhgGjuAafvxUHy4P0"
	action_type := "1"
	url := baseUrl  + "token=" + token + "&video_id=" + video_id + "&action_type=" + action_type

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
	assert.Empty(t, err)
	favor := service.GetFavoriteSrv()
	err = json.Unmarshal(body, &favor)
	assert.Empty(t, err)
	fmt.Printf(string(body))

}

//重复点赞
func TestDuplicateFavoriteAction(t *testing.T) {
	baseUrl :="http://192.168.30.128:4000/douyin/favorite/action?"
	video_id := "1"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ0ZXN0MyIsInBhc3N3b3JkIjoiZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2UiLCJhdmF0YXIiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9hdmF0YXIuanBnIiwiYmFja2dyb3VuZF9pbWFnZSI6Imh0dHA6Ly9yejJuODd5Y2suaG4tYmt0LmNsb3VkZG4uY29tL2JhY2tncm91bmQucG5nIiwic2lnbmF0dXJlIjoi5Y-I5p2l55yL5oiR55qE5Li76aG15ZWmfiIsImNyZWF0ZWRfYXQiOiIyMDIzLTA4LTI5VDIzOjI5OjA1LjY1Ni0wNzowMCIsImV4cCI6MTY5MzM4MDU0NSwiaWF0IjoxNjkzMzc2OTQ1LCJpc3MiOiJ0ZXN0MyJ9.hZxgHYARp_xs9hCylBav9YYwFqkhgGjuAafvxUHy4P0"
	action_type := "1"
	url := baseUrl  + "token=" + token + "&video_id=" + video_id + "&action_type=" + action_type

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
	assert.Equal(t, strings.Contains(str,"重复点赞"),true)
}