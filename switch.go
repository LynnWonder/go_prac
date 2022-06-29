package main
import "fmt"

func switch1(name string) {
	// 省略若干条语句
	switch name {
	case "Golang":
		fmt.Println("A programming language from Google.")
	case "Rust":
		fmt.Println("A programming language from Mozilla.")
	default:
		fmt.Println("Unknown!")
	}
}

func switch2 () {
	names := []string{"Golang", "Java", "Rust", "C"}
	switch name := names[0]; name {
	case "Golang":
		fmt.Println("=====>")
		fmt.Println("A programming language from Google.")
		// fallthrough 用来强制执行下一个 case 代码块，注意判断类型的时候不能用（如下面的函数）
		fallthrough
	case "Rust":
		fmt.Println("=====>")
		fmt.Println("A programming language from Mozilla.")
	default:
		fmt.Println("=====>")
		fmt.Println("Unknown!")
	}
}

/**
 switch语句。它与一般形式有两点差别。
 第一点，紧随case关键字的不是表达式，而是类型说明符。
 类型说明符由若干个类型字面量组成，且多个类型字面量之间由英文逗号分隔。
 第二点，它的switch表达式是非常特殊的。
 这种特殊的表达式也起到了类型断言的作用，但其表现形式很特殊，如：v.(type)，其中v必须代表一个接口类型的值。注意，该类表达式只能出现在类型switch语句中，且只能充当switch表达式。
 */
func switch3()  {
	v := 11
	switch i := interface{}(v).(type) {
	case int, int8, int16, int32, int64:
		fmt.Printf("A signed integer: %d. The type is %T. \n", i, i)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("A unsigned integer: %d. The type is %T. \n", i, i)
	default:
		fmt.Println("Unknown!")
	}
}

func main () {
	switch2()
	switch3()
}