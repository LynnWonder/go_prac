package main

import (
	"fmt"
	"time"
)

func addAsync(a int, b int) chan int {
	// 使用管道接收结果，注意需要设置一个缓冲位，否则没有取结果的话这个 goroutine 会被阻塞
	resultChan := make(chan int, 1)
	go func() {
		// 在新的 goroutine 中计算结果，并将结果发送到管道
		resultChan <- a + b
	}()
	return resultChan
}

func addWithCallback(a int, b int, callback func(sum int)) {
	go func() {
		// 在新的 goroutine 中计算结果，并将结果传递给回调函数
		sum := a + b
		callback(sum)
	}()
}

func main() {
	// 方法的两个参数
	a := 4
	b := 5

	// 阻塞式异步，如要调试请注释掉下面非阻塞式异步的内容
	// 从管道中接收结果，这一步是阻塞的，因为在等待结果的产出
	sum := <-addAsync(a, b)
	fmt.Println(sum)

	// 非阻塞式异步，是采用回调函数的方式
	// 调用方法的时候加上回调函数
	// 这个回调函数会在得到结果之后执行
	addWithCallback(a, b, func(sum int) {
		fmt.Println(sum)
	})
	// 防止 main goroutine 比异步任务的 goroutine 先退出
	time.Sleep(time.Second)
}

