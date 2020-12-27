package main

import (
    "fmt"
    "log"
    "github.com/gobuffalo/packr"
    )

const (
    Learner string = "Bob"
    StartLearnTime string = "2020-12-12"
)

// 申明并定义变量
var fristWord string 

func main(){
    box := packr.NewBox("./templates")

  s, err := box.FindString("admin/index.html")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(s)
    // 打印未初始化全局变量
    fmt.Println(fristWord)
    // 打印 全局初始化后变量值
    fristWord = "Hello,Golang!!\n"
    fmt.Printf("初始化后的值fristWord   %s",fristWord)
}
