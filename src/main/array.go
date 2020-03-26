/*
*golang 基本数据类型 数组
 */
package main

import "fmt"

const (
	MONDAY = iota + 1
	QUESDAY
	WENSDAY
	STURSDAY
	FRIDAY
)

var (
	// 声明数组
	rates1 [5]float64
	// 声明数组扑克牌,不初始化
	pukesArr [54]string
	// 声明话语,不初始化
	// words  [...] string
	// 数组初始化,指定长度和初始字面值
	countrys = [4]string{"USA", "德国", "意大利", "china"}
	// 不指定长度 使用[...],长度由后面的字面值的个数决定
	// citys str s
	// 打折折算率
	rates = [...]float64{.83, 0.95, 0.78, 0.82, 0.89}
)

// 获取打折金额
func getRatePrice(WeekDay int, Price float64) float64 {
	payment := 0.0
	// var WeekDay int = 0
	switch WeekDay {
	case MONDAY:
		payment = rates[0] * Price
	case QUESDAY:
		payment = rates[1] * Price
	case WENSDAY:
		payment = rates[2] * Price
	case STURSDAY:
		payment = rates[3] * Price
	case FRIDAY:
		payment = rates[4] * Price
	default:
		return 0.0
	}
	return payment
}

// 打印数组
func TestArr() {
	fmt.Println(rates1, pukesArr)
	fmt.Println(countrys)
	fmt.Println(rates)
	fmt.Println(MONDAY, QUESDAY, WENSDAY, STURSDAY, FRIDAY)
}

func main() {
	TestArr()
	price := getRatePrice(QUESDAY, 208)
	fmt.Print(price)
}
