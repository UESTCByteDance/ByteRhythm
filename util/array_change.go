package util

import "strconv"

// StringArray2IntArray 将字符串数组转化为整数数组
func StringArray2IntArray(strArray []string) []int {
	var intArray []int
	for _, str := range strArray {
		num, _ := strconv.Atoi(str)
		intArray = append(intArray, num)
	}
	return intArray
}
