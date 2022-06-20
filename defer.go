// defer 特性
//    1. 关键字 defer 用于注册延迟调用。
//    2. 这些调用直到 return 前才被执行。因此，可以用来做资源清理。
//    3. 多个 defer 语句，按先进后出的方式执行。
//    4. defer 语句中的变量，在 defer 声明时就决定了。Each time a "defer" statement executes,
//    the function value and parameters to the call are evaluated as usualand saved anew
//    but the actual function is not invoked.

package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}


// 延迟调用参数在注册时求值或复制，可用指针或闭包延迟读取
func test() {
	x,y :=10,20

	defer func(i int) {
		// y 被闭包引用
		println("defer:", i, y)
	}(x) // x 被复制

	x += 10
	y += 100
	println("x=", x, "y=", y)
}
func main() {
	//var whatever [5]struct{}
	//for i := range whatever {
	//	defer fmt.Println(i)
	//}
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		// 这里也可以佐证 defer 后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行函数
		//defer t.Close()

		// 要想改变上述现象可以用这种方式
		t2 := t
		defer t2.Close()
	}

	test()
}