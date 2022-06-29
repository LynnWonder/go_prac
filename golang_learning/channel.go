package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道
func download0(url string) {
	time.Sleep(time.Second*2)
	fmt.Println("start to download", url)
	time.Sleep(time.Second*2)
	// 将 url 发送给信道
	ch <- url
}

func main() {
	for i :=0; i < 3; i++ {
		var buffer bytes.Buffer
		buffer.WriteString("a.com/")
		buffer.WriteString("-")
		buffer.WriteString(strconv.Itoa(i))
		go download0(buffer.String())
	}
	for i:=0;i<3;i++{
		// tip 只有数据被 receiver 处理的时候 sender 才会阻塞，即此时阻塞等待并发协程返回信息
		msg := <-ch
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}