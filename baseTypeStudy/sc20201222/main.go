/**
  学习如何使用golang 操作文件
*/
package main

import (
    "fmt"
    "os"
    "time"
    "log"
    //"bufio"
)

// 逐行写入文件
func WriteFile(filepath string,strs []string) {
    file,err := os.OpenFile(filepath,os.O_CREATE|os.O_RDWR|os.O_TRUNC,0755)
    if err != nil {
       fmt.Printf("文件打开失败错误:%v\n",err)
    }
    for index,str := range strs {
        fmt.Printf("index=%d,value=%v\n",index,str)
    	// file.Write([]byte(str+"\n"))
    	file.WriteString(str+"\n")
    }
    defer file.Close()
}

// 使用bufio 读写文件
func WriteFileByBufio(srcFile,destFile string,strs []string){
    // 获取源文件的大小的详细信息
    // 判断源文件是否存在
    srcfileinfo,err := os.Stat(srcFile)
    if err != nil {
        fmt.Printf("文件不存在%v\n",srcFile)
        os.Exit(2)
    }else {
        fmt.Printf("文件%v存在,且文件属性如下:-----------\n",srcFile)
        srcfilename := srcfileinfo.Name()
        srcfilesize := srcfileinfo.Size()
        srcfilemode := srcfileinfo.Mode()
        srcfilemodtime := srcfileinfo.ModTime()
        srcfileisDir := srcfileinfo.IsDir()
        fmt.Printf("文件名:%T,%v\n",srcfilename,srcfilename)
        fmt.Printf("文件大小:%T,%v\n",srcfilesize,srcfilesize)
        fmt.Printf("文件权限位:%T,%v\n",srcfilemode,srcfilemode)
        fmt.Printf("文件修改时间:%T,%v\n",srcfilemodtime,srcfilemodtime)
        if srcfileisDir {
        	fmt.Printf("%v 是一个目录\n",srcfilename)
        }else{
        	fmt.Printf("%v 不是一个目录,是一个文件\n",srcfilename)
        }
    } 
    srcfile,err := os.Open(srcFile)
    if err != nil {
        fmt.Printf("源文件文件打开或创建失败:%v",err)
        os.Exit(1)
    }
    // 预装buf
    var rbc,rbc1 int
    var buf_err error
    var data string
    for {
        buf_data := make([]byte,1024)
        rbc1,buf_err = srcfile.Read(buf_data)
        if buf_err != nil {
            errInfo := buf_err.Error()
            if errInfo == "EOF" {
                break
            }else {
                log.Fatalf("Fatal error:%v\n",buf_err)
            }
        }
        rbc += rbc1
        data += string(buf_data)
    }
    fmt.Printf("一共读取%d字节,读取到的文件内容是\n----%s----\n",rbc,string(data))
    currentTime := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("当前时间是%T,%v..... 开始使用bufio读写文件\n",currentTime,currentTime)   
}

// 使用ioutil 读写文件


// 拷贝文件

func main() {
    hostname,_ := os.Hostname()
    currentdir,_ := os.Getwd()
    //parentdir,_ := os.Getwd()
    envs := os.Environ()
    fmt.Printf("操作系统的主机名称是:%s\n",hostname)
    fmt.Printf("当前的目录是:%s\n",currentdir)
    fmt.Printf("操作系统环境变量的类型是:%T\n",envs)
    fmt.Printf("%v\n",envs[0])
    //fmt.Printfn
    WriteFile("/tmp/aaa.txt",envs)
    WriteFileByBufio("/tmp/aaa.txt","/tmp/bbb.txt",[]string{"aaaaa"})
}


