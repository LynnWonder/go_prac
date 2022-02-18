package main

import "fmt"

func main() {
	// if 没有括号
	if a :=19; a<20 {
		fmt.Println("age is", a)
	} else {
		fmt.Println("age is too short")
	}

	choose(1)
}

func choose(a int) {
	// 太奇怪了，假如 fallthrough 了，那么相当于之后的内容都会执行
	// 结论： fallthrough 之后的内容不会再进行判断下一条表达式结果是否为 true
	switch a {
	case 1:
		fmt.Println("男")
		// 如果想继续往下执行那么不是用 continue 而是 fallthrough
		fallthrough
	case 2:
		fmt.Println("女")
		fallthrough
	default:
		fmt.Println("人")
	}
}