// 关于 golang 中的 json Marshal and Unmarshal
package main

import (
	"encoding/json"
	"fmt"
)

// tip
//  可导出成员：首字母大写的变量
type Stu struct {
	Name string `json:"name"`
	Age int
	Gender bool
	sex string // sex 不是一个可导出成员
	Class *Class `json:"class"`
}

type Class struct {
	Name string
	Grade int
}

func main() {
	// 此时返回的是指向新内存的指针
	//stu = new(Stu)
	//stu.Name = "Jerry"
	//stu.Age = 1
	//stu.Gender = True
	// 此时，返回的是值
	stu := Stu {
		Name: "Jerry",
		Age:1,
		Gender: true,
		sex: "female",
	}
	cla := new(Class)
	cla.Name = "test"
	cla.Grade = 1
	stu.Class = cla

	// tip
	//  https://pkg.go.dev/encoding/json#Marshal
	//  Marshal 返回的是 []byte 如果想要方便查看需要 string() 转化一下
	//  只要是可导出成员（变量首字母大写），都可以转成json。因成员变量sex是不可导出的，故无法转成json。

	// tip Marshal
	//  如果变量打上了json标签，如Name旁边的 `json:"name"` ，那么转化成的json key 就用该标签“name”，否则取变量名作为key，如“Age”，“High”。
	//  bool类型也是可以直接转换为json的value值。Channel， complex 以及函数不能被编码json字符串。当然，循环的数据结构也不行，它会导致 marshal 陷入死循环。
	//  指针变量，编码时自动转换为它所指向的值，如cla变量。
	// （当然，不传指针，Stu struct的成员Class如果换成Class struct类型，效果也是一模一样的。
	//  只不过指针更快，且能节省内存空间。）
	jsonStu,err := json.Marshal(stu)
	if err != nil {
		fmt.Println("stu 生成 json 字符串错误")
	}
	fmt.Printf("=====>检查数据类型， %v, %v\n", stu, cla)
	fmt.Printf("=====>查看 json 字符串， %v\n", string(jsonStu))
}