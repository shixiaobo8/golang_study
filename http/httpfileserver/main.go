/**
 http文件目录浏览器
**/
package main

import (
    "net/http"
    "log"
    "flag"
    "fmt"
    "strconv"
    "net"
    "io/ioutil"
    "encoding/json"
    "strings"
)


type ipInfo struct {
  Ip string `ip`
  //City string `city`
  //Region string `region`
  //Country string `country`
  //Loc string `loc`
  //Org string `org`
  //Timezone string `timezone`
  //Readme string `readme`
}



func checkArgs()(httpPort int,fileServerDir,urlPrefix,domain string){
    // step 1: define variable
    // step 2: point vairable and  default value and usage info
    port := flag.Int("port",19993,"httpfileserver端口号")
    httpPort = *port
    flag.StringVar(&fileServerDir,"dir","/tmp","httpfileserver文件服务器目录")
    flag.StringVar(&urlPrefix,"urlprefix","/files/","http访问url前缀")
    flag.StringVar(&domain,"domain","blog.devops90.cn","域名")
    // step 3:
    flag.Parse()
    //args := flag.Args()
    //fmt.Printf("%v\n",args)
    // 获取到的参数个数解析
    //fmt.Printf("一共获取到%d个参数\n",len(args)) 
    //if len(args) != 0 {
    //    for index,arg := range args {
    //        fmt.Printf("第%d个参数:%s\n",index,arg)
    //    }
    //}
    //fmt.Printf("%v  %v\n",httpPort,fileServerDir)
    return
}

func getPublicIp() (ip string){
    resp,err := http.Get("http://ipinfo.io");if err != nil{
        log.Fatalf("获取公网ip失败:%s",err)
    }
    defer resp.Body.Close()
    ipres,_ := ioutil.ReadAll(resp.Body)
    //ips,err := net.LookupIP("blog.devops90.cn");if err != nil{
    //    fmt.Printf("%s",err)
    //}
    //fmt.Printf("%T  %v\n",ipres,ipres)
    ips := string(ipres)
    //ips := "[" + string(ipres) + "]"
    //fmt.Printf("%T\n",ips)
    //fmt.Printf("%s\n",ips)
    //ipinfo := []ipInfo{}
    ipinfo := make(map[string]string)
    err1 := json.Unmarshal([]byte(ips),&ipinfo);if err1 != nil{
        log.Fatal("json解析错误!%v",err1)
    }
    ip = ipinfo["ip"]
    //ip = ipinfo[0].Ip
    //fmt.Printf("%s22222\n",ipinfo)
    //fmt.Printf("%s\n",ip)
    return 
}

func getInetIps() (ips []string) {
    inetAddrs,err := net.InterfaceAddrs();if err != nil {
        fmt.Printf("获取网卡地址失败\n%s",err)
    }
    //for index,addr :=range inetAddrs{
    for _,addr :=range inetAddrs{
        ip := net.ParseIP(strings.Split(addr.String(),"/")[0])
        // fmt.Printf("%d,%s\n",index,ip)
        // 判处ipv6
        if strings.Contains(ip.String(),"."){
            ips = append(ips,ip.String())
        }
    }
    return 
}

func main() {
    publicIp := getPublicIp()
    inetIps := getInetIps()
    //fmt.Printf("%v\n",inetIps)
    port,dir,urlprefix,domain := checkArgs()
    addr := ":" + strconv.Itoa(port)
    httpFileServerDir := http.Dir(dir)
    fileHandler := http.FileServer(httpFileServerDir)
    stripPrefixfileHandler := http.StripPrefix(urlprefix,fileHandler)
    httpServer := http.Server{Addr:addr,Handler:stripPrefixfileHandler}
    listenHosts := make([]string,2+len(inetIps),(2+cap(inetIps))*2)
    listenHosts = append(listenHosts,domain)
    listenHosts = append(listenHosts,publicIp)
    //listenHosts := []string{publicIp,domain}
    //fmt.Printf("lis %d,%d\n",len(listenHosts),cap(listenHosts))
    //fmt.Printf("inet %d,%d\n",len(inetIps),cap(inetIps))
    copy(listenHosts,inetIps)
    //fmt.Printf("lis %d,%d\n",len(listenHosts),cap(listenHosts))
    //fmt.Printf("%v\n",listenHosts)
    fmt.Printf("http文件服务器启动成功,请尝试如下地址访问:\n")
    for _,listenHost := range listenHosts{
        fmt.Printf("%s\t","http://"+listenHost+addr+urlprefix)
    }
    log.Fatalf("http文件服务器启动失败:%v",httpServer.ListenAndServe())	
}
