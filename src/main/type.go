/**
golang 基本数据类型之 结构体： 将gelang 中的各种基本数据类型结构组合在一起 构成类似 其它语言的 oop 中的class 和 继承
**/
package main

import (
	"fmt"
)

const (
	TODAY  = "20200401"
	FEMALE = iota
	MALE
)

var (
	status string = "愚人节快乐!"
)

// 人类
type person struct {
	name string
	age  int
	sex  int
}

// 职员
type staff struct {
	person
	salary     int
	department string
}

// 公司
type company struct {
	companyName    string
	companyPositon string
	companyMajor   string
	companyIncome  int
	staffs         []staff
}

// 老板
type boss struct {
	company
	person
}

// 定义职员打工赚钱的方法
func StaffWork(sf staff) (salary int) {
	return sf.salary*14 + salary*130
}

// 定义老板经营公司的方法
func BossManageCompany(bos boss) (bonus int) {
	// 获取boss 的所有员工工资
	allStaffsSalarys := 0
	for _, staff := range bos.staffs {
		allStaffsSalarys += StaffWork(staff)
	}
	// 获取一年的经营额 = 公司收入
	others := 2000000
	profit := bos.companyIncome - allStaffsSalarys - others
	return profit
}

func main() {
	// 定义一个人
	Bob := person{name: "Bob", age: 30, sex: MALE}
	Eddie := person{name: "Eddie", age: 30, sex: MALE}
	Shuwen := person{name: "Eddie", age: 30, sex: FEMALE}
	staffBob := staff{
		person:     Bob,
		salary:     10000,
		department: "技术部运维组",
	}
	staffEddie := staff{
		person:     Eddie,
		salary:     12000,
		department: "技术部运维组",
	}
	staffMage := staff{
		person:     Shuwen,
		salary:     20000,
		department: "人事部",
	}
	roy := person{
		name: "roy",
		age:  42,
		sex:  MALE,
	}
	staffs := []staff{staffBob, staffEddie, staffMage}
	weicheche := company{"深圳市喂车科技", "深圳市南山区粤海街道软件产业基地X栋Y楼", "智慧油站行业", 6e8, staffs}
	fmt.Println(Bob)
	fmt.Println(weicheche)
	weichecheBoss := boss{
		company: weicheche,
		person:  roy,
	}
	fmt.Println(weichecheBoss)
	//for key,value := range weichecheBoss {
	//	fmt.Println(key,value)
	//}
	// 获取公司老板一年的赚取的利润
	fmt.Println(weichecheBoss.companyPositon+weichecheBoss.companyName+"的老板"+weichecheBoss.name+"经营"+weichecheBoss.companyMajor+"一年赚取的收入为:", BossManageCompany(weichecheBoss))
}
