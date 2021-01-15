package main

import "fmt"

/**
函数类型的字面量由关键字func、由圆括号包裹参数声明列表、空格以及可以由圆括号包裹的结果声明列表组成。
其中，参数声明列表中的单个参数声明之间是由英文逗号分隔的。每个参数声明由参数名称、空格和参数类型组成。
参数声明列表中的参数名称是可以被统一省略的。结果声明列表的编写方式与此相同。结果声明列表中的结果名称也是可以被统一省略的。
并且，在只有一个无名称的结果声明时还可以省略括号。


函数类型的零值是 nil,这意味着一个未被显shi赋值的函数类型的变量必为 nil
 */

// 结果声明是带名称的
func myFunc(part1 string, part2 string) (result string) {
	result = part1 + part2
	return
}

func myFunc1(part1 string, part2 string) string {
	return part1 + part2
}

var val =  func(part1 string, part2 string) string {
	return part1 + part2
}

var result = func(part1 string, part2 string) string {
	return part1 + part2
}("1", "2")


func main () {
	fmt.Print(result)
}

