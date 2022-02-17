// 声明 main.go 所在的包，go 通过包来组织代码，一个文件夹即一个包，包内暴露类型或方法供其他包使用
package main

// fmt 是 go 的一个标准库/包，用来处理标准输入输出
import "fmt"

// main 函数是整个程序的入口，main 函数所在的包名必须为 main
func main() {
	var a int8 = 10
	var b float32 = 11.1
	var c byte = 'c'
	str1 := "golang"
	str2 := "go123"
	// 直接输出
	fmt.Println("hello world", a, b, c)
	// 格式化输出
	fmt.Printf("%s %s\n", str1, str2)
}

