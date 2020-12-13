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

func ReturnPointAddr(a, b int) *int {
	sum := a + b
	return &sum
}

// 计算斐波那契数列
func fid(n int) []int {
	x, y := 0, 1
	farr := []int{x}
	for i := 0; i < n; i++ {
		x, y = y, x+y
		farr = append(farr, x)
	}
	return farr
}

func ReturnPointAddrValue(a, b int) int {
	return *ReturnPointAddr(a, b)
}

func newInt() *int {
	return new(int)
}

// 测试new 不是一个关键字  而是一个内置函数
func delta(old, new int) int {
	return new - old
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

	saddr := ReturnPointAddr(3, 6)
	fmt.Println(saddr)
	// 函数内调用的指针地址 每次都不一样
	fmt.Println(ReturnPointAddr(1, 3) == ReturnPointAddr(1, 3))
	fmt.Println(ReturnPointAddrValue(4, 6))
	fmt.Println(p1)
	fmt.Println(newInt())
	// 比较指针地址值
	p3 := newInt()
	p4 := newInt()
	p5 := p3
	// 指针地址值不通
	fmt.Println(p3 == p4)
	// 指针地址值相同
	fmt.Println(p5 == p3)
	v := delta(3, 0)
	fmt.Println(v)
	// 打印第一百个斐波那契数列的值
	fidArr := fid(20)
	fmt.Println(len(fidArr), fidArr)
}
