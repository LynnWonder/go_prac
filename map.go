package main

import "fmt"

func main()  {
	// map[K]T K 必须是可以比较的类型因此不能是切片字典或者函数
	mm :=map[string]string{"a":"a", "b":"b", "c":"c"}
	b :=mm["b"]
	fmt.Println("b===>", b)
	mm["b"]= b + "2"
	fmt.Printf("mm===> %v\n", mm)
	mm["d"]=""
	d := mm["d"]
	e := mm["e"]
	// 在Go语言中有这样一项规定，
	//即：对于字典值来说，如果其中不存在索引表达式欲取出的键值对，
	//那么就以它的值类型的空值（或称默认值）作为该索引表达式的求值结果。
	//由于字符串类型的空值为""，所以mm[5]的求值结果即为""。
	fmt.Printf("d: %v，e: %v\n", d,e)
	e, ok:= mm["e"]
	// 针对字典的索引表达式可以有两个求值结果。
	// 第二个求值结果是bool类型的。
	// 它用于表明字典值中是否存在指定的键值对。在上例中，变量ok必为false。因为mm中不存在以5为键的键值对
	fmt.Printf("e: %v，ok: %v\n", e,ok)
	// 删除键值对
	delete (mm, "d")
	// 有则删除，无则不
	fmt.Printf("after deleting 'd''===> %v\n", mm)
}