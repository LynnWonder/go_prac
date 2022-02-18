package main

// tip golang 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发。
// 另一个例子参考 channel
import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 例如我们希望并发下载 N 个资源，多个并发协程之间不需要通信，那么就可以使用 sync.WaitGroup，等待所有并发协程执行结束。
func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second*2)
	// 减去一个计数
	wg.Done()
}

func main() {
	for i :=0; i < 3; i++ {
		wg.Add(1)
		var buffer bytes.Buffer
		buffer.WriteString("a.com/")
		buffer.WriteString("-")
		buffer.WriteString(strconv.Itoa(i))
		go download(buffer.String())
	}
	wg.Wait()
	fmt.Println("Done!")
}