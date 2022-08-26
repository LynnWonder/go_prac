package main

import (
	"fmt"
)

type B struct {
	thing int
}

type C struct {
	thinga *string
	thingb *string
}

// change 方法接受一个指向 B 的指针并改变它的内部成员
func (b *B) change() { b.thing = 1 }

// write 方法通过拷贝接受 B 的值并输出 B 的内容
func (b B) write() string { return fmt.Sprint(b) }

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

// tip
//  指针方法和值方法都可以在指针或者非指针上调用
func main() {
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	// b2 指向的是一个指向新的已分配内存的指针
	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())

	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("输出 lst 相关内容===> %v (len: %d)\n", lst, lst.Len())

	// 指针

	plst := new(List)
	plst.Append(2)
	// tip 当调用 Len() 方法时指针会自动解引用，变成 (*plst).Len()
	fmt.Printf("输出 plst 相关内容===> %v (len: %d)\n", plst, plst.Len())

	// TIP 使用指针作为值，可以避免结构体默认赋初值，从而判断哪些参数传入，哪些没传
	str := "string"
	test := C{
		thinga: &str,
	}
	fmt.Println(test.thinga, test.thingb)
}

/* 输出：
{1}
{1}
*/
