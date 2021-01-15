package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}


// 空接口类型即是不包含任何方法声明的接口类型，用interface{}表示，常简称为空接口
// 注意在 go 文件中大写就允许在其他文件中进行访问
type Animal interface {
	Grow()
	Move(string) string
}

func (person *Person) Grow() {
	person.Age++
}

func (person *Person) Move(add string) string{
	temp :=person.Address
	person.Address = add
	return temp
}

func main() {
	p := Person{"Tom", "Male", 33, "Beijing"}
	// 表达式&p（&是取址操作符）的求值结果是一个*Person类型的值，即p的指针。
	v := interface{}(&p)

	h, ok := v.(Animal)

	fmt.Printf("val of h: %v, ok is %v", h, ok)
}
