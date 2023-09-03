package wrapper

import "github.com/afex/hystrix-go/hystrix"

var RegisterFuseConfig = hystrix.CommandConfig{
	Timeout:                1000,
	RequestVolumeThreshold: 5000, // 10秒内的请求量，默认值是20，如果超过20那么就判断是否熔断
	ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
	SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	MaxConcurrentRequests:  50000,
}

var LoginFuseConfig = hystrix.CommandConfig{
	Timeout:                1000,
	RequestVolumeThreshold: 5000, // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
	ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
	SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	MaxConcurrentRequests:  50000,
}
var UserInfoFuseConfig = hystrix.CommandConfig{
	Timeout:                1000,
	RequestVolumeThreshold: 5000, // 熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
	ErrorPercentThreshold:  50,   // 错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
	SleepWindow:            5000, // 过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	MaxConcurrentRequests:  50000,
}
