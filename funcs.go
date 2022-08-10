package global

import "strconv"

// SliceIntToStr
// 将整型数组转换成字符串数组
func SliceIntToStr(intSlice []int) (strSlice []string) {
	strSlice = make([]string, len(intSlice))
	for i := range intSlice {
		strSlice[i] = strconv.Itoa(intSlice[i])
	}
	return
}
