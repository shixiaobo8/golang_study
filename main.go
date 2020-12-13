package main

import (
    "fmt"
    )

const (
    Learner string = "Bob"
    StartLearnTime string = "2020-12-12"
)

// 申明并定义变量
var fristWord string 

func main(){
    // 打印未初始化全局变量
    fmt.Println(fristWord)
    // 打印 全局初始化后变量值
    fristWord = "Hello,Golang!!\n"
    fmt.Printf("初始化后的值fristWord   %s",fristWord)
}
