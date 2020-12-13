/**
golang 基本数据类型rune类型(int32的类型别名)
*/
package main

import (
	"fmt"
)

var (
	rune1 rune = 'x'
	rune2 int  = 'A'
	r1    int  = '你'
	r2    int  = '好'
)

func main() {
	fmt.Printf("%U is\n", rune2)
	fmt.Printf("%c is\n", rune1)
	fmt.Printf("%d")
}
