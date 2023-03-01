package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Id   string
}

func main() {
	o := &User{"Kevin", "k3v"}
	var b interface{}

	/**
	  如果是指针类型，有三种方式判断类型
	*/
	// 方法 1
	b = o
	if bUser, ok := b.(*User); ok {
		fmt.Println("方法一：b是一个User类型", bUser.Name)
	}

	// 方法 2
	switch b.(type) {
	case *User:
		fmt.Println("方法二：b是一个User类型")
	}

	// 方法 3
	ty := reflect.TypeOf(b)
	fmt.Println("方法三：", ty, "ty.Name()获取到的是空值")
	// 报错， panic: reflect: NumField of non-struct type
	// fmt.Println("方法三：", "ty.NumField()获取到的是空值", ty.NumField())

	/**
	b 如果是非指针类型，有三种方式判断类型
	*/
	// 给b赋值的不是一个指针 是一个普通对象
	b = User{"K3vin", "kev"}
	if bUser, ok := b.(User); ok {
		fmt.Println("方法一：b是一个User类型", bUser.Name)
	}
	switch b.(type) {
	case User:
		fmt.Println("方法二：b是一个User类型")
	}

	ty1 := reflect.TypeOf(b)
	fmt.Println("方法三：", ty1, "ty1.Name()获取到的不再是空值", ty1.Name())
}
