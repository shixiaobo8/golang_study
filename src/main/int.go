/**
golang 基本数据类型之 int 类型
---------------------------------------------
byte:   字节. uint 的别名
rune:   int32 的别名
int:    有符号整型,
int8:   有符号整型,占用8bit,数值范围: -127~127 (-2的7次方-1~2的7次方-1)
int16:  有符号整型,占用16bit,数值范围:
int32:  有符号整型,占用32bit,数值范围:
int64:  有符号整型,占用64bit,数值范围:
uint:   无符号整型:
uint8:  无符号整型,占用8bit,数值范围: 0~256
uint16: 无符号整型,占用16bit,数值范围: 0~65536
uint32: 无符号整型,占用32bit,数值范围: 0~2的32次方
uint64: 无符号整型,占用64bit,数值范围: 0~2的64次方
-------------------------------------------------
 */

package main

import (
	f1 "fmt"
	"reflect"
)

// 定义常量
const (
	PI = 3.141592653
	X = 2
	Y = 3
	Z,W = 4,5
)

// 定义变量 基本数据类型
var (
	// 字符串
	s1 string = "global variable str 1"
	// byte
	// int
	s2 int = 10243452343245333
	// byte 等同于int8，常用来处理ascii字符
	sn3 uint8 = 255
	s3 int8 = 127
	//
	s4 int16 = 25945
	s5 int64 = 102434523432453
	// bool 类型
	s6 bool = false
	// int32 的别名 runa
	s7 int32 = 2324213
	// 8 进制int
	s8 int = 0500
	// 16 进制 int
	s16 int = 0xBadFace3
	// 浮点类型 float32
	ff1 float32 = 0.
	ff2 float32 = .034
	ff3 float32 = 3223E8
	ff4 float64 = 003.3332
	ff5 float64 = 08.32E-11
	)

// 定义 一般类型

// 定义结构体

// 定义接口


func main()  {
	cwd_start := "========常量start========"
	f1.Println(cwd_start)
	f1.Println(reflect.TypeOf(PI),PI)
	f1.Println(reflect.TypeOf(X),X)
	f1.Println(reflect.TypeOf(Y),Y)
	f1.Println(reflect.TypeOf(Z),Z,reflect.TypeOf(W),W)
	cwd_end := "========常量end========"
	f1.Println(cwd_end)
	vwd_start := "========变量start========="
	vwd_end := "========变量end========="
	f1.Println(vwd_start)
	str := "Hello, Golang22!"
	f1.Println(reflect.TypeOf(str),str)
	f1.Println(reflect.TypeOf(s2),s2)
	f1.Println(reflect.TypeOf(ff5),ff5)
	f1.Println(reflect.TypeOf(ff4),ff4)
	f1.Println(reflect.TypeOf(ff2),ff2)
	f1.Println(reflect.TypeOf(ff3),ff3)
	f1.Println(reflect.TypeOf(s7),s7)
	f1.Println(reflect.TypeOf(s16),s16)
	f1.Println(vwd_end)
	// test1()
	// testFunc()
}