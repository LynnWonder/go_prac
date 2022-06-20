package main

import "fmt"
import "github.com/LynnWonder/go_prac/myself"

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

/**
TIP 在默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
 注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。
 而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
 注意2：map、slice、chan、指针、interface默认以引用的方式传递。
 */

// TIP 不定长参数通过在类型前加 ... 标识
func summary (n ...int) int{
	res :=0
	for _, val :=range n {
		res = res + val
	}
	return res
}

func add (x,y int) (s int) {
	s = x  + y
	// 如果函数体内没有重新命名 s 那么直接像下面这样隐式返回就可以
	return
}

// 和 defer 结合
func addD(x, y int)(z int) {
	defer func() {
		println(z)
	}()
	z = x + y
	return z+200
}


func main () {
	fmt.Print(result)
	fmt.Println("===计算和==>", summary(1,2,3), summary(0,1))
	// tip 注意使用 slice 对象做变参的时候，必须展开
	sl := []int{1,3,4}
	res := summary(sl...)
	fmt.Println("=== res is ===>", res)
	// TIP 结合 defer 函数之后的执行顺序：z=z+200 -> call defer -> return
	println("=== 结合 defer===>",addD(1,2))

	// init 函数总是早于 main 函数执行
	var test = myself.Test
	println("===myself.test===", test)
}

