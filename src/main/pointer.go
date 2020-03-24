/**
golang 基本数据类型之 指针类型
go语言不支持指针运算
函数中允许返回局部变量的地址
*/
package main

import (
	"fmt"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
)

var (
	str string = "hello 你好"
	p1         = &str
)

// 结构体
type (
	User struct {
		name string
		age  int
	}
)

func RetunPointAddr(a, b int) *int {
	sum := a + b
	return &sum
}

func main() {
	fmt.Println(B, KB, MB, GB)
	fmt.Println(str)
	fmt.Println(p1)
	fmt.Println(*p1)

	//
	bob := User{
		name: "Bob",
		age:  30,
	}

	// 指针结构体
	p := &bob
	fmt.Println(bob.age)
	fmt.Println(p.name)

	saddr := RetunPointAddr(3, 6)
	fmt.Println(saddr)
}
