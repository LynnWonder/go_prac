package main

import "fmt"

func main()  {
	a :=add0(1,2)
	fmt.Println("======>add0 计算结果",a)

	fmt.Println("======>错误处理", get(20))
	fmt.Println("======>不抛出错误", get(2))
}

// 使用关键字 func，参数可以有多个，返回值也支持有多个。
// 特别地，package main 中的func main() 约定为可执行程序的入口
// 注意一个包里不能有重名的函数
func add0(num1 int, num2 int) int {
	return num1 + num2
}


// 处理错误，使用 defer and recover
func get(index int) (res int) {
	// 协程退出前会执行 defer 挂载的任务，因此如果触发了 panic 那么控制权就交给了 defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			res = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}