package util

import (
	"strconv"
	"strings"
)

// StringArray2IntArray 将字符串数组转化为整数数组
func StringArray2IntArray(strArray []string) []int {
	var intArray []int
	for _, str := range strArray {
		num, _ := strconv.Atoi(str[strings.Index(str, ":")+1:])
		intArray = append(intArray, num)
	}
	return intArray
}
