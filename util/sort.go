package util

import (
	"sort"
	"strconv"
	"strings"
)

func SortKeys(keys []string) []string {
	// 对key按照第二个%d排序（%d可能不止一位），并改为降序排序
	sort.Slice(keys, func(i, j int) bool {
		// 提取第二个%d后的数字进行比较
		num1, _ := strconv.Atoi(keys[i][strings.Index(keys[i], ":")+1:])
		num2, _ := strconv.Atoi(keys[j][strings.Index(keys[j], ":")+1:])
		return num1 > num2
	})

	// 只保留前30个视频
	if len(keys) > 30 {
		keys = keys[:30]
	}

	return keys
}
