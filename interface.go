package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}


// tip
// 	空接口类型即是不包含任何方法声明的接口类型，用 interface{} 表示，常简称为空接口
// 	任何其他类型都实现了空接口，因此它常常可以作为函数的参数以表示任意类型
// 注意在 go 文件中大写就允许在其他文件中进行访问
type Animal interface {
	Grow()
	Move(string) string
}

func (person *Person) Grow() {
	person.Age++
}

// tip 一个接收者是 Person 结构体的方法 Move
// 	实现了 Animal 接口
func (person *Person) Move(add string) string{
	temp :=person.Address
	person.Address = add
	return temp
}

func main() {
	p := Person{"Tom", "Male", 33, "Beijing"}

	var animal Animal
	// 如果 p 直接是通过 new 创建的那么就不用 & 去取指针了
	animal = &p
	// tip
	//  这就是 golang 中的多态，即同一种类型在不同的实例上能够表现出不同的行为
	fmt.Printf("接口变量 animal 包含一个指向 p 变量的引用，通过它可以调用 Grow 方法，%v\n", animal.Move("test"))
	// 表达式&p（&是取址操作符）的求值结果是一个*Person类型的值，即p的指针。
	v := interface{}(&p)

	h, ok := v.(Animal)

	fmt.Printf("val of h: %v, ok is %v", h, ok)
}
