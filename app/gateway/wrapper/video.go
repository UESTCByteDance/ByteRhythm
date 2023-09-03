package wrapper

import (
	"ByteRhythm/idl/video/videoPb"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
)

var FeedFuseConfig = hystrix.CommandConfig{
	Timeout:                1000,
	RequestVolumeThreshold: 5000, // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
	ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
	SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	MaxConcurrentRequests:  50000,
}

var PublishFuseConfig = hystrix.CommandConfig{
	Timeout:                5000,
	RequestVolumeThreshold: 5000, // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
	ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
	SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	MaxConcurrentRequests:  50000,
}
var PublishListFuseConfig = hystrix.CommandConfig{
	Timeout:                1000,
	RequestVolumeThreshold: 5000, // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
	ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
	SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	MaxConcurrentRequests:  50000,
}

func NewVideo(vid int64, title string) *videoPb.Video {
	return &videoPb.Video{
		Id:            vid,
		Author:        nil,
		Title:         title,
		PlayUrl:       "",
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
}

// DefaultFeed 降级函数
func DefaultFeed(res interface{}) {
	models := make([]*videoPb.Video, 0)
	for i := 0; i < 30; i++ {
		models = append(models, NewVideo(int64(i), "降级视频流"+strconv.Itoa(i)))
	}
	result := res.(*videoPb.FeedResponse)
	result.VideoList = models
}
