// main.go 所在包 Go 语言中使用包来组织代码。一般一个文件夹即一个包，包内可以暴露类型或方法供其他包使用。
package main

// import “fmt”：fmt 是 Go 语言的一个标准库/包，用来处理标准输入输出。
import "fmt"



// main 函数是整个程序的入口，main 函数所在的包名也必须为 main。
func main() {
	fmt.Println("Hello golang World")
	// 值类型
	a :=2
	b :=0
	b=a
	a = 1
	// tip 果然，a 发生变化后 b 并没有随之改变，
	//   因为是在内存中对 a 的值进行了拷贝，而导致 a 发生了变化之后，b 并不会随之发生变化
	fmt.Printf("当 a 发生变化时, b 是否发生变化呢===>%d, %d\n", a, b)
	// tip 由此可见 a 和 b 的内存地址并不一样
	fmt.Printf("看一下 a 和 b 的内存地址===>%p, %p\n", &a, &b)
}
