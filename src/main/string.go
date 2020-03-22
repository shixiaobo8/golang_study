/**
golang 基本数据类型之 string 字符串
1. 字符串 可以转化为  byte数组 或者 rune(int32) 数组
2. 字符串可以做slice 切片操作,但是不可以改变数组元素的值
3. 字符串的长度不一定是字符的长度,使用len获取的不是真正的字符串长度
 */
package main

import (
	"fmt"
)

const (
	VERSION = 1.0
)

var (
	fristWord string = "Hello 你好 golang!"
)

func main(){
	fmt.Println(fristWord)
	// 字符串为一个常量不能修改
	fristWord := "Hello modify str"
	fmt.Println(fristWord)
	// 通过切片索引或许字节元素
	subFristWord := fristWord[3:13]
	fmt.Println(subFristWord)
	// 字符串切片后仍为 字符串
	fmt.Println(fristWord[0:8])
	// 转换为unicode 字数组
	strToUnicodeArr := []rune(fristWord)
	fmt.Println(strToUnicodeArr)
	// 转换为字节数组
	strToByteArr := []byte(fristWord)
	fmt.Println(strToByteArr)
	// 字符串拼接
	str1 := "你好"
	str2 := "World"
	hello := str1 + str2
	fmt.Println(hello)
	// 获取字符串长度 这里的中文不是两个字节,是一个rune(int32) unicode字节码
	fmt.Println(len(hello))
	fmt.Println(len([]rune(str1)))
	// 遍历字符串字节byte数组
	for i:=0;i<len(hello);i++{
		fmt.Println(hello[i])
	}
	// 遍历字符串rune 数组
	for i:=0;i<len([]rune(hello));i++{
		fmt.Println(hello[i])
	}
}