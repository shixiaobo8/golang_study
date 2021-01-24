// 包名
package main

// 导入包
import (
	"fmt"
	"io/ioutil"
	"log"

	// "log"
	// "encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// 定义常量
const APOLLO_CONFIG_SERVER_PORT = "80"

// 定义全局变量
var apollo_config_server_base_url string = "http://apollo.wcc.cn:" + APOLLO_CONFIG_SERVER_PORT
var configListUrl string = apollo_config_server_base_url + "/apps?appIds="

// 一般类型声明
// type newType aaa

// 结构体声明
// type gopher struct{}

// 接口声明
// type golang interface{}

// 函数
func getLoginCookie() (jar http.CookieJar, httpcookiejarerror error) {
	// apollo登陆url
	apollo_login_url := apollo_config_server_base_url + "/signin"
	fmt.Printf("%s\n", apollo_login_url)

	// cookie
	//cookiejar, err := http.CookieJar{}
	jar, httpcookiejarerror = cookiejar.New(nil)
	fmt.Printf("%v", jar)
	if httpcookiejarerror != nil {
		fmt.Printf("cookie jar init false!")
		return
	}
	client := http.Client{Jar: jar}
	username := "apollo"
	password := "juqhN64FQCQ3q"
	loginSubmit := "登录"
	formValues := url.Values{}
	formValues.Add("username", username)
	formValues.Add("password", password)
	formValues.Add("login-submit", loginSubmit)
	fmt.Printf("formData:%s\n", formValues.Encode())

	request, err := http.NewRequest("POST", apollo_login_url, ioutil.NopCloser(strings.NewReader(formValues.Encode())))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(request)
	// 构造post from data
	//loginData := "username=" + username + "&password=" + password + "&login-submit=%E7%99%BB%E5%BD%95"

	// 构造http request 客户端
	// client, err := http.Post(apollo_config_server_base_url, login_data)
	// if err != nil {
	// 	log.Fatalf("client err%s", err.Error())
	// }
	// request, _ := http.NewRequest("POST", apollo_login_url, strings.NewReader(login_data))
	// 访问前的cookie
	// fmt.Printf("before request cookie is%v", request.Cookies())
	// 提交post请求并接收数据

	// resp, err := http.Post(apollo_login_url, "application/x-www-form-urlencoded", strings.NewReader(loginData))
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
		fmt.Println(jar)
	}
	defer resp.Body.Close()
	return
	// fmt.Println(login_url)
}

// 获取列表数据
func getApolloConfigList(jar http.CookieJar) {
	httpClient := &http.Client{Jar: jar}
	fmt.Printf("%v\n", httpClient.Jar)
	httpreq, err := http.NewRequest("GET", configListUrl, nil)
	httpreq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	httpresp, err := httpClient.Do(httpreq)
	if err != nil {
		fmt.Printf("%s", "请求失败"+err.Error())
	} else {
		resBody, _ := ioutil.ReadAll(httpresp.Body)
		fmt.Printf("%s", string(resBody))
	}
}

// main 函数
func main() {
	// fmt.Println("Hello,Apollo!")
	jar, _ := getLoginCookie()
	getApolloConfigList(jar)

}
