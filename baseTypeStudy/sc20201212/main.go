/**
-----------------------------------------
	golang 基本数据类型练习 20201212
        带颜色输出的不适用于windows
-----------------------------------------
*/

package main

import (
    "fmt"
    "strconv"
    "bufio"
    "strings"
    //"unsafe"
    "os"
    "math"
)

// 常量
const (
    INT_INIT = iota   // 初始化int 暂用内存字节数
    INT8
    INT16
    INT32
    INT64 
)

// 变量 初始化并赋值
var BaseTypeKeyWord string = `
    golang 7 大基本数据类型如下:
                        1. 布尔型: bool 默认值为false
                        2. 整型: 默认值int  根据操作系统位数决定  
			3. 浮点型:  float32 float64
			4. 复数类型:   complex64 complex128
			5. 字符串:	string
			6. 字符: rune
			7. 错误类型: error
`

// 声明全局变量不初始化
var (
    BoolKeyWord,IntKeyWord,FloatKeyWord,ComplexKeyWord,StringKeyWord,ByteKeyWord,ErrorKeyWord string
)

// int最小字节数
var (
    MinIntByte = 1
    MinIntBit = 8
)


//  package 文件中的init 编译执行的第一部分
func init() {
    fmt.Printf("今天是20201212,开始练习...\n")
    fmt.Printf(BaseTypeKeyWord+"\n")
}

// bool 的使用
func TestBoolUseAge(){
    var isFlag bool
    fmt.Printf("isFlag default values is %#v\n",isFlag)
    BoolNotice := `注意点: 和其他语言不一样,非0不是true,golang中的int不能和bool作比较?
    testFlag := 1
    fmt.Printf(bool(testFlag) == isFlag)  // 这种是错误的`
    fmt.Printf(BoolNotice+"\n")
    // string 转化为bool
    var boolstring string = "true"
    newBool,_ := strconv.ParseBool(boolstring)
    fmt.Printf("before stringtoBool : %T and boolstring=%v,after conv : %T and newBool=%v \n",boolstring,boolstring,newBool,newBool)
    // bool 转换为string
    newBoolString := fmt.Sprintf("%t",newBool)
    fmt.Printf("bool convert to string:%T,%v\n",newBoolString,newBoolString)
    fmt.Printf("\""+"%s"+"\"\n",newBoolString)
    fmt.Printf("\""+"%q"+"\"\n",newBoolString)
    
}

// 整型的使用方法 switch 版本
func TestIntUseAgeBySwitch() {
    //  对全局变量IntKeyWord赋值 
    IntKeyWord = `
                        整型:
                             int: 根据操作系统的位数决定,32位操作系统,就相当于int32,64位操作系统就相当于int64
                             int8:  占用内存1字节(1Byte=8bit) 数值范围: - 2*7-1 到 2*7-1
			     int16: 占用内存2字节(2Byte=16bit) 数值范围: - 2*15-1 到 2*15-1
			     int32(rune): 占用内存4字节(4Byte=32bit) 数值范围: - 2*31-1 到 2*31-1
			     int64: 占用内存8字节(8Byte=64bit) 数值范围: - 2*63-1 到 2*63-1
			     uint8(byte): 无符号 占用内存1字节  数值范围:  0 到 2*7-1
			     uint16: 无符号 占用内存2字节  数值范围:  0 到 2*15-1
			     uint32: 无符号 占用内存4字节  数值范围:  0 到 2*31-1
			     uint64: 无符号 占用内存8字节  数值范围:  0 到 2*63-1
`
    fmt.Printf(IntKeyWord+"\n")
    scannerInfo := ` 请输入一个int类型的的名称,将打印出该整型的占用内存字节数,以及数值的范围:
         例如:  输入8或int8 将输出如下结果:
                             int8:  占用内存1字节(1Byte=8bit) 数值范围: - 2*7-1 到 2*7-1
         如输入异常讲会提示重新输入
请输入int类型: `
    scanner := bufio.NewScanner(os.Stdin)  
    fmt.Printf("\033[36;33;3m %s \033[0m",scannerInfo)
    for scanner.Scan() {
        st := scanner.Text()
	switch {
	    // 判断输入并对输入的字符串进行逻辑处理
        case st == "int8" || strings.Contains(st,"int8"):
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(8)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        case st == "int16" || strings.Contains(st,"int16"):
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(16)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        case st == "int32" || strings.Contains(st,"int32"):
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(32)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        case st == "int64" || strings.Contains(st,"int64"):
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(64)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        default:
            inputErrInfo := "未检测到输入的有效位数,请重新输入!"
            fmt.Printf("\033[31;31;5m %s \033[0m\n",inputErrInfo)
            fmt.Printf("\033[36;33;3m %s \033[0m",scannerInfo)
        }
    }
}


// 整型的使用方法 if else 版本
func TestIntUseAge() {
    //  对全局变量IntKeyWord赋值 
    IntKeyWord = `
                        整型:
                             int: 根据操作系统的位数决定,32位操作系统,就相当于int32,64位操作系统就相当于int64
                             int8:  占用内存1字节(1Byte=8bit) 数值范围: - 2*7-1 到 2*7-1
			     int16: 占用内存2字节(2Byte=16bit) 数值范围: - 2*15-1 到 2*15-1
			     int32(rune): 占用内存4字节(4Byte=32bit) 数值范围: - 2*31-1 到 2*31-1
			     int64: 占用内存8字节(8Byte=64bit) 数值范围: - 2*63-1 到 2*63-1
			     uint8(byte): 无符号 占用内存1字节  数值范围:  0 到 2*7-1
			     uint16: 无符号 占用内存2字节  数值范围:  0 到 2*15-1
			     uint32: 无符号 占用内存4字节  数值范围:  0 到 2*31-1
			     uint64: 无符号 占用内存8字节  数值范围:  0 到 2*63-1
`
    fmt.Printf(IntKeyWord+"\n")
    scannerInfo := ` 请输入一个int类型的的名称,将打印出该整型的占用内存字节数,以及数值的范围:
         例如:  输入8或int8 将输出如下结果:
                             int8:  占用内存1字节(1Byte=8bit) 数值范围: - 2*7-1 到 2*7-1
         如输入异常讲会提示重新输入
请输入int类型: `
    scanner := bufio.NewScanner(os.Stdin)  
    fmt.Printf("\033[36;33;3m %s \033[0m",scannerInfo)
    for scanner.Scan() {
        st := scanner.Text()
        /**
         ** 字符串与切片的演练
        st_arr := st[0:]
        fmt.Printf("%T,%v\n",st_arr,st_arr)
        // 遍历字节数组1
        for i:=0;i<len(st_arr);i++ {
            fmt.Println(i,st_arr[i])
        }
        // 遍历rune数组
        for index,value := range st_arr {
            fmt.Println(index,value)
        }
        */
        // 判断输入并对输入的字符串进行逻辑处理
        if st == "int8" || strings.Contains(st,"int8") {
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(8)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        } else if st == "int16" || strings.Contains(st,"int16") {
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(16)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        } else if st == "int32" || strings.Contains(st,"int32") {
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(32)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        } else if st == "int64" || strings.Contains(st,"int64") {
            fmt.Printf("获取到您输入的int类型:%s\n",st)
            mem,maxValueRange := PrintIntMemRange(64)
            fmt.Printf("%s的占用内存字节数%d byte,其取值范围是: -%d到%d\n",st,mem,maxValueRange,maxValueRange)
        } else {
            inputErrInfo := "未检测到输入的有效位数,请重新输入!"
            fmt.Printf("\033[31;31;5m %s \033[0m\n",inputErrInfo)
            fmt.Printf("\033[36;33;3m %s \033[0m",scannerInfo)
        }
        // 无需转换格式了
        //stInt,err := strconv.ParseInt(st,10,8)
	//if err != nil {
           // fmt.Println(err)
            // panic(err)
        //}
        // 对输入的字符串进行逻辑处理
        // fmt.Printf("%T,%q\n",st,st)
        // fmt.Printf("af: %T,%v\n",stInt,stInt)
        
    }
}

// 打印int整型的内存占用字节数以及数值范围
func PrintIntMemRange(intbit int) (mem int,maxValueRange int64) {
	mem = intbit/8
        maxValueRange = int64((math.Pow(float64(2),float64((intbit-1))))-1)
        return 
}

// 或者命令行参数

func main() {
   TestBoolUseAge()
   // TestIntUseAge()
   // switch case 版本
   TestIntUseAgeBySwitch()
}
