package main

import "fmt"

func main() {
	// 这是标准声明
	var a int8 = 10
	var b float32 = 12.2
	var c1 byte = 'a'
	var msg = "Hello World"
	// TIP 这是简短声明，注意简短声明只能在函数内部使用
	ok := false
	ok = true

	fmt.Println("this is a", a)
    fmt.Println("this is b", b)
    fmt.Println("this is c1", c1)
	fmt.Println("this is msg", msg)
	fmt.Println("this is ok", ok)
}
