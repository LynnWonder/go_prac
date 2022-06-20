// golang 中也有匿名函数，且匿名函数可以赋值给变量，作为结构字段或者在 channel 里传送

package main

func main() {
	fn := func() {
		println("这是一个匿名函数 ==>Hello world!")
	}
	fn()

	// 匿名函数切片
	fns := []func(x int) int{
		func(x int) int {return x+1},
		func(x int) int {return x+2},
	}
	println("匿名函数切片==>", fns[0](100))

	// 匿名函数作为结构体的值
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello world"},
	}
	println(d.fn())

	// channel of function
	fc :=make(chan func() string,2)
	fc <- func() string {return "Hello World!"}
	println((<-fc)())
}