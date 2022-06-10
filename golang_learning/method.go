package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

type User struct {
	Name string
	Email string
}
func (u *User) Notify(param string)(s string){
	return u.Name + "'s email is" + u.Email + param
}
// tip golang 方法总是绑定对象实例，并隐式将实例作为第一实参即 receiver
//  func (recevier type) methodName(参数列表)(返回值列表){}

// 定义一个结构体上的简单方法，实现相加的功能
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a+tn.b+param
}

// TIP 结构体和方法的区别
//  1. 对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。
//  2. 对于方法（如 struct 的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。

type IntVector []int
// 定义一个非结构体上的简单方法
func (V IntVector) Sum(s int) int {
	res :=0
	for _,x:=range V {
		// tip _ 指的是键值，在数组中指的就是索引
		res +=x
	}
	return s + res
}
func main() {
	two1 :=&TwoInts{1,2}
	two2 :=&IntVector{7,2,3}
	user :=&User{"lynn", "test@gmail.com"}

	fmt.Println("使用结构体类型上的方法=====>", two1.AddToParam(3))
	fmt.Println("使用非结构体类型上的方法=====>", two2.Sum(3))
	fmt.Println("=======>", user.Notify("~~"))
}