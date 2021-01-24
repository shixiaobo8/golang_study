// 包名
package main

// 导入包
import (
	"fmt"
	"io/ioutil"
	"log"

	// "log"
	"net/http"
	"strings"
)

// 定义常量
const APOLLO_CONFIG_SERVER_PORT = "8070"

// 定义全局变量
var apollo_config_server_base_url string = "http://192.168.85.203:" + APOLLO_CONFIG_SERVER_PORT

// 一般类型声明
// type newType aaa

// 结构体声明
// type gopher struct{}

// 接口声明
// type golang interface{}

// 函数
func getLoginCookie() {
	username := "apollo"
	password := "fwadmin123"

	// apollo登陆url
	apollo_login_url := apollo_config_server_base_url + "/signin"

	// 构造post from data
	loginData := "username=" + username + "&password=" + password + "&login-submit=%E7%99%BB%E5%BD%95"

	// 构造http request 客户端
	// client, err := http.Post(apollo_config_server_base_url, login_data)
	// if err != nil {
	// 	log.Fatalf("client err%s", err.Error())
	// }
	// request, _ := http.NewRequest("POST", apollo_login_url, strings.NewReader(login_data))
	// 访问前的cookie
	// fmt.Printf("before request cookie is%v", request.Cookies())
	// 提交post请求并接收数据
	fmt.Printf("%s\n", apollo_login_url)
	resp, err := http.Post(apollo_login_url, "application/x-www-form-urlencoded", strings.NewReader(loginData))
	if err != nil {
		fmt.Printf("post请求出错!", err)
	} else {
		fmt.Printf("%s\n", resp.Status)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("未接收到数据!")
		} else {
			fmt.Printf("%v", string(body))
		}
		fmt.Println(resp.Cookies())
	}
	defer resp.Body.Close()
	// fmt.Println(login_url)
}

// main 函数
func main() {
	// fmt.Println("Hello,Apollo!")
	getLoginCookie()
}

