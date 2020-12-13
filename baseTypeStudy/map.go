/**
golang 基本数据类型之map
知识点回顾:
	map 的几种定义方法：
	key: 必须为可hash的对象
	value：
	map 的操作:
	map 遍历
	map key 的遍历
	map value 的遍历
	map 是无序的,需要遍历可以先对key进行排序,然后再遍历
	map key 和value 的反转
*/
package main

import "fmt"

const (
	// 学习日期
	STUDYDAY = 20200331
)

var (
	// 今日学习心态
	studyStatus string = "is not very good!"
	// map 定义方式一：
	m1 map[int]string = map[int]string{-1: "yesterday", 1: "today", 2: "tomorrow"}
	// map 定义方式二:
	m2 = map[string]int{"not ready!": -1, "ready": 0, "doing": 1, "done": 2}
	// map 定义方式三:
	m3 = make(map[int]string)
)

// 打印一下变量^_^
func PrintMap() {
	fmt.Println(m1)
	// 获取map value
	fmt.Println("修改m1的值", m1[0])
	// 修改map value
	fmt.Println(m2)
	fmt.Println(m3)
}

// 遍历 m1 map key
func PrintMapM1Elements() {
	// 循环遍历key
	// 重新赋值之前的指针变量
	pk := &m1
	// 打印重新赋值之前的指针地址值
	fmt.Println(&pk)
	for k, _ := range m1 {
		fmt.Println(k, pk, m1[k])
		// 重新对值进行设置
		m1[k] = "好好学习golang,每天进步一点点!"
		fmt.Println(k, m1[k])
	}
	pk1 := &m1
	fmt.Println(&pk1, m1)
}

func main() {
	PrintMap()
	PrintMapM1Elements()
	fmt.Println(m1)
	// 删除map 中的值
	t1 := &m2
	fmt.Println(m2, &t1)
	delete(m2, "not ready!")
	fmt.Println(m2, &t1)
	// 添加key value
	m2["is not ok"] = -2
	m2["is ok"] = 3
	fmt.Println(m2, &t1)
	m3[200] = "ok"
	m3[300] = "rewrite"
	m3[400] = "not found"
	m3[500] = "server errors"
	fmt.Println(m3)
	// 判断key 是否存在
	s, ok1 := m3[503]
	s2, ok2 := m3[500]
	if ok1 {
		fmt.Println("存在key"+"503", s)
	} else {
		fmt.Println("不在在key503", s)
	}
	if ok2 {
		fmt.Println("存在key500", s2)
	} else {
		fmt.Println("不存在key500", s2)
	}
}
