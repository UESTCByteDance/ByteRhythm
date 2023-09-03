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

// -----------------/douyin/favorite/action/接口测试--------------------
// 点赞成功
func TestFavoriteAction(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/favorite/action/?"
	video_id := "1"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJuYW1lIjoidGVzdDIiLCJwYXNzd29yZCI6IjIwMmNiOTYyYWM1OTA3NWI5NjRiMDcxNTJkMjM0YjcwIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0zMFQwNDo1NDo1Mi0wNzowMCIsImV4cCI6MTY5MzM5NjYyMiwiaWF0IjoxNjkzMzk2NTAyLCJpc3MiOiJ0ZXN0MiJ9.6CiMJkVJDk-nSpl28qgQXmTkgGqIobQff-Zg7MBMioc"
	action_type := "1"
	url := baseUrl + "token=" + token + "&video_id=" + video_id + "&action_type=" + action_type

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

// 重复点赞
func TestDuplicateFavoriteAction(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/favorite/action/?"
	video_id := "3"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJuYW1lIjoidGVzdDIiLCJwYXNzd29yZCI6IjIwMmNiOTYyYWM1OTA3NWI5NjRiMDcxNTJkMjM0YjcwIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0zMFQwNDo1NDo1Mi0wNzowMCIsImV4cCI6MTY5MzM5NjYyMiwiaWF0IjoxNjkzMzk2NTAyLCJpc3MiOiJ0ZXN0MiJ9.6CiMJkVJDk-nSpl28qgQXmTkgGqIobQff-Zg7MBMioc"
	action_type := "1"
	url := baseUrl + "token=" + token + "&video_id=" + video_id + "&action_type=" + action_type

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
	assert.Equal(t, strings.Contains(str, "重复点赞"), true)
}

// 取消点赞
func TestCancelFavoriteAction(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/favorite/action/?"
	video_id := "3"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJuYW1lIjoidGVzdDIiLCJwYXNzd29yZCI6IjIwMmNiOTYyYWM1OTA3NWI5NjRiMDcxNTJkMjM0YjcwIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0zMFQwNDo1NDo1Mi0wNzowMCIsImV4cCI6MTY5MzM5NjYyMiwiaWF0IjoxNjkzMzk2NTAyLCJpc3MiOiJ0ZXN0MiJ9.6CiMJkVJDk-nSpl28qgQXmTkgGqIobQff-Zg7MBMioc"
	action_type := "2"
	url := baseUrl + "token=" + token + "&video_id=" + video_id + "&action_type=" + action_type

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

// -----------------/douyin/favorite/list/接口测试--------------------
func TestFavoriteList(t *testing.T) {
	baseUrl := "http://192.168.256.128:8080/douyin/favorite/list/?"
	user_id := "11"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJuYW1lIjoidGVzdDIiLCJwYXNzd29yZCI6IjIwMmNiOTYyYWM1OTA3NWI5NjRiMDcxNTJkMjM0YjcwIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0zMFQwNDo1NDo1Mi0wNzowMCIsImV4cCI6MTY5MzM5NjYyMiwiaWF0IjoxNjkzMzk2NTAyLCJpc3MiOiJ0ZXN0MiJ9.6CiMJkVJDk-nSpl28qgQXmTkgGqIobQff-Zg7MBMioc"
	url := baseUrl + "user_id=" + user_id + "token=" + token

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
	assert.Equal(t, strings.Contains(str, "获取喜欢列表成功"), true)
}
