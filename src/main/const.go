/**
golang 标识符 常量 const
常量 const 存放内存地址的内存是不可变的
 */
package main

import "fmt"

// 类似枚举的iota
const (
	c0 = iota
	c1 = iota
	c2 = iota
)

func main(){
	fmt.Println(c0,c1,c2)
}

