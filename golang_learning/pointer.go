package main

import "fmt"

func main() {
	str :="golang"
	// tip 关于指针中的 * 和 &
	// 	指针类型定义时使用 *，使用 & 获取该类型的地址：一句话来解释的话就是假如有变量 aa, 则 &aa 它的类型就是 *string
	//  取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	// p 是一个指向 str 的指针
	var p *string = &str
	fmt.Println(p)
	pp := *p
	fmt.Println(pp)

	// 改变指针 p 所指的值，自然 str 也就发生了变化
	*p = "hello world"
	fmt.Println("str====>",str)

	// 测试指针的使用
	num :=1
	add(num)
	fmt.Println("num=====>",num)
	realNum :=1
	realAdd(&realNum)
	fmt.Println("realNum=====>", realNum)
	res :=realAdd0(0)
	fmt.Println("realNum0=====>", res)


	// TIP make 和 new 的区别：golang 用 new 和 make 来为引用类型分配内存，其中 new 不太常用
	//  但 make 只能用于 slice map channel 的内存创建
}

// 一般来说，指针通常在函数传递参数，或者给某个类型定义新的方法时使用。
// Go 语言中，参数是按值传递的，如果不使用指针，函数内部将会拷贝一份参数的副本，对参数的修改并不会影响到外部变量的值。
// 如果参数使用指针，对参数的传递将会影响到外部变量。
// 综上所述，如果想影响外部变量那么就用指针
func add(param int) {
	param +=1
}

func realAdd(param *int) {
	*param +=1
}

func realAdd0 (param int) int {
	param++
	return param
}