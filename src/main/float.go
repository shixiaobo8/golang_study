/**
golang 基本数据类型 float 练习  float32/float64
1. float 的表示方法:
	1)可以使用.0023 省去小数点前面的操作或者用0填充
	2) 使用科学计数法表示:  如02.002E22 / 23.99e23
2. 一般使用 float64，精度更准确
3. float运算一般使用math库
 */
package main

import (
	f1 "fmt"
	"math"
	)

const (
	PI1 = 03.14162
)

var (
	pi float32 = 003.14156253
	pai float64 = 003.141592653e100
	)

func main ()  {
	f1.Println(PI1)
	f1.Printf("2%f\n",math.Pi)
	f1.Println(pi)
	f1.Println(pai)
	test1()
}