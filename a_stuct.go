// 这里讲解结构体匿名字段

package main

import (
	"fmt"
	"reflect"
)

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous field 结构体匿名字段，类型就是字段的名字
	innerS //anonymous field 结构体可以内嵌结构体
}


type TagType struct {
	field1 bool `important answer`
	field2 string
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)


	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

	tt := TagType{true, "Barak Obama"}
	// tip 结构体带标签后，可以使用 reflect 包来使用它获取 tag
	fmt.Printf("tt 的类型====>%v，%v,%v\n", reflect.TypeOf(tt),reflect.TypeOf(tt).Field(0),reflect.TypeOf(tt).Field(0).Tag)
}