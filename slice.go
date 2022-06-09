package main

import "fmt"

// tip 声明切片的几种方式
func main() {
	// 1.声明切片
	// tip 数组定义：
	//  var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。
	//  一旦定义，长度不能变。
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}

	// 2.:=
	s2 := []int{}

	// 3.make()，make([]int, len, cap)，若省略 cap 则 cap=len
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	// 数组也可以这样赋值，只不过会添上长度或者 ...
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)


	s := []int{0, 1, 2, 3}
	p := &s[2] // *int, tip 获取底层数组元素指针。
	*p += 100 // 对指针所指的位置的值增加 100
	// TODO 了解指针

	fmt.Println("===s 变化后===",s)
}