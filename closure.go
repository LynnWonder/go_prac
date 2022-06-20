// TIP 闭包：其实可以通俗的理解为函数里有另一个函数，即函数和与其相关的引用环境组合而成的实体。因此闭包可以用来完成信息隐藏
//   不同的引用环境和相同的函数组合可以产生不同的实例
package main

// func a() (func() int) 这样就很容易理解了
func a() func() int {
	i := 0
	b := func() int {
		i++
		println("====>",i)
		return i
	}
	return b
}

func addC(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// 返回两个闭包
func twoClosure(base int) (func(int) int, func(int) int){
	add1 := func(i int) int {
		base +=i
		return base
	}
	add2 :=func(i int) int {
		base -=i
		return base
	}
	return add1, add2
}

// golang 递归函数
func fib(i int)int{
	if i==0 {
		return 0
	}
	if i==1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func main() {
	c :=a()
	c()
	c()
	c()

	a()

	tmp1 := addC(10)
	println("===> tmp1(1)", tmp1(1))
	println("===> tmp1(2)", tmp1(2))

	tmp2 := addC(100)
	println("===> tmp2(1)", tmp2(1))
	println("===> tmp2(2)", tmp2(2))

	f1, f2 :=twoClosure(100)
	ff1, ff2 :=twoClosure(10)
	println("===> f1", f1(1))
	println("===> f2", f2(1))
	println("===> ff1", ff1(1))
	println("===> ff2", ff2(1))

	for i:=0; i<10; i++ {
		println("===>fib", i, fib(i))
	}
}