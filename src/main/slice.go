/*
golang 基本数据类型之  slice 切片
*/
package main

// 四季
const (
	SPRING = iota + 1
	SUMMER
	FALL
	WINTER
)

// 季度月份
var (
	// season
	seasonNames = [4]string{"spring", "summer", "fall", "winter"}
	// 一年十二个月份
	year    = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	seasons = make([]int, 4, 4)
	// seasonMap
	seasonMap = make(map[string][]int)
)

func PrintSeason() []int {
	return seasons
}

func printYear() [12]int {
	return year
}
