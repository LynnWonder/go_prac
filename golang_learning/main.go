// tip 声明 main.go 所在的包，go 通过包来组织代码，一个文件夹即一个包，包内暴露类型或方法供其他包使用
// tip 一般来说，golang 中一个文件夹可以作为一个 package，一个 package 中的变量、类型、方法等定义都可以互相被看到
// tip golang 也有 public 和 private 的概念，粒度是包，如果类型、接口、方法、函数、字段的首字母大写，那就是 public 的
// 反之，如果首字母为小写，那么就是 private 的，对其他包而言是不可见的
package main

// fmt 是 go 的一个标准库/包，用来处理标准输入输出
import (
	"fmt"
	"reflect"
)

// main 函数是整个程序的入口，main 函数所在的包名必须为 main
func main() {
	var a int8 = 10
	var b float32 = 11.1
	var c byte = 'c'
	str1 := "golang"
	str2 := "go语言"
	runeArr := []rune(str2)
	// 直接输出
	fmt.Println("hello world", a, b, c)
	// 格式化输出
	fmt.Printf("%s %s\n", str1, str2)
	// 获取数据类型
	fmt.Println(reflect.TypeOf(str2).Kind())
	// 可以发现字符串是以 byte 数组的形式存储的，类型是 unit8，占一个 byte,不转换的话打印的就是编码值
	fmt.Println(str2[0],reflect.TypeOf(str2[0]).Kind(), string(str2[0]))
	fmt.Println(str2[2],reflect.TypeOf(str2[2]).Kind(), string(str2[2]))
	// 通过上面的测试发现对于汉字不会通过 string(vars) 转换
	// 转换成 []rune 类型后，字符串中的每个字符，无论占多少个字节都用 int32 来表示
	fmt.Println(runeArr[2],reflect.TypeOf(runeArr[2]).Kind(), string(runeArr[2]))
}

