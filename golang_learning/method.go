package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

// 定义一个结构体上的简单方法，实现相加的功能

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a+tn.b+param
}

type IntVector []int
// 定义一个非结构体上的简单方法
func (V IntVector) Sum(s int) int {
	res :=0
	for _,x:=range V {
		// tip _ 指的是键值，在数组中指的就是索引
		res +=x
	}
	return s + res
}
func main() {
	two1 :=&TwoInts{1,2}
	two2 :=&IntVector{7,2,3}

	fmt.Println("使用结构体类型上的方法=====>", two1.AddToParam(3))
	fmt.Println("使用非结构体类型上的方法=====>", two2.Sum(3))

}