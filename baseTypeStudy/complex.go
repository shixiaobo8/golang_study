/*
golang 基本数据类型之 复数类型 complex,有complex64 和complex128 两种类型,复数结构如:   实部(float32)+虚部(float32)i
可以使用两种方式定义complex:
1. 使用内置函数complex定义
	var c1 complex128 = complex(实部float位,虚部float位i)
2. 使用 变量定义法定义
	var c2 complex64 = .0232+09.00i
3. 获取复数complex 实部的方法real
	var shibu float32 = real(c2)
4. 获取复数complex 虚部的方法imag ,注意这里不是image哦
	var xubu float43 = imag(c2)
 */
package main

import (
	"fmt"
	)
const (
	PI = 3.141592
)

var (
	c1 complex64 = 3.1 + 5i
	c2 complex128 = 3.1 + 6i
)

func main (){
	fmt.Println(PI)
	fmt.Println(c1)
	fmt.Println(c2)
	// 构造一个复数
	v := complex(2.1e11,3.1e10)
	fmt.Println(v)
	// 返回复数实部
	fmt.Println(real(v))
	// 返回复数虚部
	fmt.Println(imag(v))
}