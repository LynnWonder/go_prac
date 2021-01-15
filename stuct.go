package main

import "fmt"

/**
结构体类型的字面量由关键字type、类型名称、关键字struct，以及由花括号包裹的若干字段声明组成。
其中，每个字段声明独占一行并由字段名称（可选）和字段类型组成。

 与代表函数值的字面量类似，我们在编写一个结构体值的字面量时不需要先拟好其类型。这样的结构体字面量被称为匿名结构体。与匿名函数类似，我们在编写匿名结构体的时候需要先写明其类型特征（包含若干字段声明），再写出它的值初始化部分。下面，我们依照结构体类型Person创建一个匿名结构体：
 */
type Personn struct {
	Name   string
	Gender string
	Age    uint8
}
// 讲指针的时候会再进行讲解


// 一个方法的接收者类型是其所属类型的指针类型而不是该类型本身，该方法则为一个指针方法。反之则为值方法
// 那么这里为什么会用指针方法而不是值方法呢
// 因为如果是用 Person 那么此时代表的是 p 值的拷贝而不是 p 的值
// 在调用 grow 方法时 go 会将 p 的值复制一份并将其作为此次调用的当前值，因此 Grow 方法中的 person.Age++ 语句的执行会使这个副本的 Age 字段的值变为 34，
// 而 p 的 Age 字段的值却依然是 33。这就是问题所在

// 而如果使用 *Person 那么此时 person代表的是p的值的指针的副本。
// 指针的副本仍会指向p的值。另外，之所以选择表达式person.Age成立，是因为如果Go语言发现person是指针并且指向的那个值有Age字段，那么就会把该表达式视为(*person).Age。
// 其实，这时的person.Age正是(*person).Age
func (person *Personn) Grow() {
	person.Age++
}
func main() {
	// 与代表函数值的字面量类似，我们在编写一个结构体值的字面量时不需要先拟好其类型。
	//这样的结构体字面量被称为匿名结构体。与匿名函数类似，我们在编写匿名结构体的时候需要先写明其类型特征（包含若干字段声明），再写出它的值初始化部分。
	//下面，我们依照结构体类型Person创建一个匿名结构体，与此同时匿名结构体是不能有方法的


	// 熟悉面向对象编程的同学可能已经意识到，包含若干字段和方法的结构体类型就相当于一个把属性和操作封装在一起的对象。
	// 不过要注意，与对象不同的是，结构体类型（以及任何类型）之间都不可能存在继承关系。
	// 实际上，在Go语言中并没有继承的概念。不过，我们可以通过在结构体类型的声明中添加匿名字段（或称嵌入类型）来模仿继承
	p := struct {
		Name   string
		Gender string
		Age    uint8
	}{"Robert", "Male", 33}
	fmt.Printf("value of p: %v\n", p)
	q :=Personn{"Tom", "Male", 30}
	q.Grow()
	fmt.Printf("value of q: %v\n", q)
}