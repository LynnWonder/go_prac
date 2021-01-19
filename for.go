package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	/**
	for语句还有另外一种编写方式，那就是用range子句替换掉for子句。
	range子句包含一个或两个迭代变量（用于与迭代出的值绑定）、特殊标记:=或=、关键字range以及range表达式。
	其中，range表达式的结果值的类型应该是能够被迭代的，包括：字符串类型、数组类型、数组的指针类型、切片类型、字典类型和通道类型。
	 */
	// todo: 为何索引不连续
	for i, v := range "Go语言" {
		fmt.Printf("%d: %c\n", i, v)
	}
}