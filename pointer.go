package main

// Pointers.
// Go has pointers.
// A pointer holds the memory address of a value.
// The type *T is a pointer to a T value. Its zero value is nil

// 这里可以看其在结构体中的使用
// 此处仅做练习题

import "fmt"

type MyInt struct {
	n int
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt *MyInt) Decrease() {
	myInt.n--
}

func main() {
	mi := MyInt{}
	mi.Increase()
	mi.Increase()
	mi.Decrease()
	mi.Decrease()
	mi.Increase()
	fmt.Printf("%v\n", mi.n == 1)
}