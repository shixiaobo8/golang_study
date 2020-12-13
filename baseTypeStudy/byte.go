/**
golang 基本数据类型之 byte 类型 字节类型
*/
package main

import (
	"fmt"
	"time"
)

const (
	STUDY_DAY = 3
	STUDY_NUM = iota
	STUDY_DAYS
	STUDY_TIME
)

// byte 暂用1个字节 8bit
var (
	b1 byte = 's'
)

func test() {
	today := time.Now().Format("2020-02-22 00:00:00")
	fmt.Println(today)
	fmt.Println(b1)
}

func main() {
	fmt.Println(STUDY_TIME, STUDY_DAYS, STUDY_NUM, STUDY_TIME)
	test()
}
