package main

import "fmt"

// TIP 变量声明：:= 的方式并不适用于全局变量声明，若要声明全局变量请使用 var
func main() {
	// 第一种声明方式
	var a int = 100
	fmt.Printf("type of a is %T\n", a)

	var b = 100
	fmt.Printf("type of b is %T\n", b)

	c := "test"
	fmt.Printf("type of c is %T\n", c)
}
