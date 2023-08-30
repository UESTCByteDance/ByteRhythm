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
	video_id := "9"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIsInVzZXJuYW1lIjoidGVzdDEiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0yOVQyMDoxOTozOC0wNzowMCIsImV4cCI6MTY5MzM3NDE2NCwiaWF0IjoxNjkzMzcwNTY0LCJpc3MiOiJ0ZXN0MSJ9.6VOX2tUA32meapWynWxzNepgXC4ZqBNHhqZ2DgA486U"
	action_type := "1"
	url := baseUrl  + "token=" + token + "&video_id" + video_id + "&action_type" + action_type

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
	video_id := "9"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIsInVzZXJuYW1lIjoidGVzdDEiLCJwYXNzd29yZCI6ImUxMGFkYzM5NDliYTU5YWJiZTU2ZTA1N2YyMGY4ODNlIiwiYXZhdGFyIjoiaHR0cDovL3J6Mm44N3ljay5obi1ia3QuY2xvdWRkbi5jb20vYXZhdGFyLmpwZyIsImJhY2tncm91bmRfaW1hZ2UiOiJodHRwOi8vcnoybjg3eWNrLmhuLWJrdC5jbG91ZGRuLmNvbS9iYWNrZ3JvdW5kLnBuZyIsInNpZ25hdHVyZSI6IuWPiOadpeeci-aIkeeahOS4u-mhteWVpn4iLCJjcmVhdGVkX2F0IjoiMjAyMy0wOC0yOVQyMDoxOTozOC0wNzowMCIsImV4cCI6MTY5MzM3NDE2NCwiaWF0IjoxNjkzMzcwNTY0LCJpc3MiOiJ0ZXN0MSJ9.6VOX2tUA32meapWynWxzNepgXC4ZqBNHhqZ2DgA486U"
	action_type := "1"
	url := baseUrl  + "token=" + token + "&video_id" + video_id + "&action_type" + action_type

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