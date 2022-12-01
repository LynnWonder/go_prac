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
	/*
		tip: reflect 实现了运行时反射，允许程序操作任意类型的数据，经常使用的场景是使用 reflect.TypeOf() 来获取数据类型
			reflect.ValueOf() 则是返回其运行时的值，indirect 返回指针类型的值指向的值
	*/
	content := "Name"
	peo := new(People)
	//Elem()获取指针对应元素的值
	v := reflect.ValueOf(peo).Elem()
	//CanSet():判断值有没有被设置，有设置:True,没有设置：false
	fmt.Println(v.FieldByName(content).CanSet())

	//需要修改属性的内容时，要求结构体中属性名首字母大写才可以设置
	v.FieldByName(content).SetString("test")
	v.FieldByName("Address").SetString("Beijing")
	peo1 := &People{
		"lynn",
		"aa@163.com",
	}
	sl := []*People{peo, peo1}
	val := reflect.ValueOf(sl)
	fmt.Println(peo)
	fmt.Println("reflect 获取值===>", val, sl, val.Index(0).Interface())
	// indirect 将获取指针指向的值，如果不是指针类型就不要回做额外的操作
	indirectVal := reflect.Indirect(reflect.ValueOf(sl))
	fmt.Println("reflect.indirect 获取值===>", indirectVal.Index(0).Interface())

	peo1Val := reflect.ValueOf(peo1)
	peo1IndirectVal := reflect.Indirect(reflect.ValueOf(peo1))
	fmt.Println("reflect 获取值 peo1Val===>", peo1Val.Interface())
	// indirect 这里获取的是指针指向的值
	fmt.Println("reflect.indirect 获取值===>", peo1IndirectVal.Interface())

	mm := map[string]string{"a": "aVal", "b": "bVal", "c": "cVal"}
	mmVal := reflect.ValueOf(mm)
	fmt.Println("reflect 获取  mmVal 值的类型===>", mmVal.Kind())
	mapRange := mmVal.MapRange()
	for mapRange.Next() {
		fmt.Println("map 中的每一个键值对===>", mapRange.Key().Interface(), mapRange.Value().Interface())
	}
	//kind := val.Kind()
	//fmt.Println("reflect 获取值的类型===>", kind == reflect.Slice)
}
