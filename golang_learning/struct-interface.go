package main

import (
	"bytes"
	"fmt"
	"reflect"
)

// tip 结构体类似于其他语言中的 class，可以在结构体中定义多个字段，为结构体实现方法，实例化等。
// tip 接下来我们定义一个结构体 Student，并为 Student 添加 name，age 字段，并实现 hello() 方法。

type Student struct {
	name string
	age int
}
func (stu *Student) hello(person string) string {
	var buffer bytes.Buffer
	buffer.WriteString("hello ")
	buffer.WriteString(person)
	buffer.WriteString(", my name is ")
	buffer.WriteString(stu.name)
	buffer.WriteString(", age is ")
	fmt.Println("age is ====>",stu.age)
	return buffer.String()
}
// tip golang 中并不显式的声明实现了哪一个接口，而是直接实现接口的方法即可
func (stu *Student) getName() string {
	return stu.name
}
func main() {
	// 实例化
	// 没有被显性赋值的变量会被赋予默认值
	stu :=&Student{
		name: "Tom",
	}
	res :=stu.hello("jack")
	fmt.Println("res is====>", res)

	// tip
	//  相当于实例化 Student 后，强制转换成接口类型 Person
	var stu0 Person=&Student{
		name: "jerry",
	}
	fmt.Println("======>实现接口", stu0.getName())

	// tip 接口也可以类型转换为实例
	pp :=stu0.(*Student)
	fmt.Println("======>接口转换为实例", pp, reflect.TypeOf(pp))
}

// tip
//  接口 interfaces 定义了一组方法的集合，
//  接口不能被实例化，一个类型可以实现多个接口
// 	golang 中，并不需要显式的声明实现了哪个接口，只需要直接实现该接口对应的方法即可

type Person interface {
	getName() string
}


