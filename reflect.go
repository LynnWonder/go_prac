package main

import (
	"fmt"
	"reflect" //反射包
)

type People struct {
	Name    string
	Address string
}

func main() {
	content := "Name"
	peo := new(People)
	//Elem()获取指针对应元素的值
	v := reflect.ValueOf(peo).Elem()
	//CanSet():判断值有没有被设置，有设置:True,没有设置：false
	fmt.Println(v.FieldByName(content).CanSet())

	//需要修改属性的内容时，要求结构体中属性名首字母大写才可以设置
	v.FieldByName(content).SetString("minger")
	v.FieldByName("Address").SetString("shenzhen")
	peo1 := &People{
		"lynn",
		"aa@163.com",
	}
	sl := []*People{peo, peo1}
	val := reflect.ValueOf(sl)
	fmt.Println(peo)
	fmt.Println("reflect 获取值===>", val.Index(0).Interface())
	kind := val.Kind()
	fmt.Println("reflect 获取值的类型===>", kind == reflect.Slice)
}
