package util

import (
	"sort"
	"strconv"
)

func SortKeys(keys []string) []string {
	sort.Slice(keys, func(i, j int) bool {
		num1, _ := strconv.Atoi(keys[i])
		num2, _ := strconv.Atoi(keys[j])
		return num1 > num2 // 改为大于号以实现降序排序
	})
	//只取30个视频
	if len(keys) > 30 {
		keys = keys[:30]
	}
	return keys
}
