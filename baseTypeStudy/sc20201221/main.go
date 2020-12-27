package main

import (
    "fmt"
    "time"
)

// 初级跑步
type runner interface {
     running()
} 

// 马拉松飞行
type flyer interface {
     running()
} 


type runnerPartner struct {
      name string
      runTime string
}

func (rp runnerPartner) running(){
      fmt.Printf("%s在%s开始跑步%s\n",rp.name,rp.runTime)
} 

func main (){
      nowTime := time.Now()
      now := nowTime.String()
      runnerYaoHong := runnerPartner {name:"耀鸿",runTime:now}
      runnerHuangPing := runnerPartner {name:"黄平",runTime:now}
      runnerBob := runnerPartner {name:"Bob",runTime:now}
      var r runner
      var f runner
      r = runnerYaoHong
      f = runnerHuangPing
      r.running()
      f.running()
      runnerBob.running() 
      var r1 runner
      r1 = runnerBob
      r1.running()
}
