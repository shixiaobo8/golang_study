/**
使用golang http 模拟登陆 爬取apollo config server 数据
 */

// 包名
package utils

// 导入包
import s  "fmt"

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 定义常量
const APOLLO_CONFIG_SERVER_PORT = "8070"

// 定义全局变量
var Apollo_config_server_base_url string = "http://192.168.85.203:" + APOLLO_CONFIG_SERVER_PORT

// 一般类型声明
// type newType aaa

// 结构体声明
// type gopher struct{}

// 接口声明
// type golang interface{}

// 测试func
func testFunc(){
	strwords := "I am i private func"
	s.Println(strwords)
}

// 函数
func LoginCookie(){
	// username := "apollo"
	// password := "fwadmin123"

	// apollo登陆url
	apollo_login_url := Apollo_config_server_base_url + "/signin"

	// 构造post from data
	login_data := `{"username":"apollo","password":"fwadmin123"}`

	// 构造http request 客户端

	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
		Timeout: time.Duration(10)*time.Second,//设置超时
	}

	request,_ := http.NewRequest("POST",apollo_login_url,strings.NewReader(login_data))

	// 提交post请求并接收数据
	resp,err := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("post请求出错!",err)
	}else{
		fmt.Println(string(body))
	}
	// defer resp.Body.Close()
	fmt.Println(apollo_login_url)
}

// main 函数
// func main(){
	// fmt.Println("Hello,Apollo!")
	// LoginCookie()
//}
