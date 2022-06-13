package main

import (
	"fmt"
	"sort"
)

func main()  {
	// tip map[K]T , K 必须是可以比较的类型因此不能是切片字典或者函数
	mm :=map[string]string{"a":"aVal", "b":"bVal", "c":"cVal"}
	b :=mm["b"]
	fmt.Println("b===>", b)
	mm["b"]= b + "2"
	fmt.Printf("mm===> %v\n", mm)
	mm["d"]=""

	d := mm["d"]
	e := mm["e"]
	// tip 在Go语言中有这样一项规定，
	//  即：对于字典值来说，如果其中不存在索引表达式欲取出的键值对，
	//  那么就以它的值类型的空值（或称默认值）作为该索引表达式的求值结果。
	//  由于字符串类型的空值为""，所以mm[d]的求值结果即为""。
	//  这样就比较坑啊
	fmt.Printf("d: %v，e: %v\n", d,e)


	// 判断某个键是否存在
	e, ok:= mm["e"]
	// 针对字典的索引表达式可以有两个求值结果。
	// 第二个求值结果是bool类型的。
	// 它用于表明字典值中是否存在指定的键值对。在上例中，变量ok必为false。因为mm中不存在以5为键的键值对
	fmt.Printf("e: %v，ok: %v\n", e,ok)


	// map 的遍历
	// tip 遍历 map 时的元素顺序与添加键值对的顺序无关
	for k := range mm {
		fmt.Println("====遍历显示 mm 的 key:", k)
	}
	// 要想按制定的顺序遍历 map，则需要提前给 key 排序
	keys := make([]string, 0, 10)
	for k:=range mm{
		keys =append(keys,k)
	}
	fmt.Println("===当前所有的 keys ===>", keys)
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("===按顺序遍历 map 的值===>", mm[key])
	}

	// 删除键值对
	delete (mm, "d")
	// 有则删除，无则不
	fmt.Printf("after deleting 'd''===> %v\n", mm)


	// 元素为 map 类型的切片
	// 指定了长度
	var mapSlice = make([]map[string]string, 3)
	for idx, val := range mapSlice {
		fmt.Printf("init index:%d value:%v\n", idx, val)
	}
	mapSlice[0] = make(map[string]string,10)
	mapSlice[0]["name"] = "test"
	for idx, val := range mapSlice {
		fmt.Printf("after index:%d value:%v\n", idx, val)
	}

	// 值为切片类型的 map
	var sliceMap = make(map[string][]string,3)
	var value = make([]string,0,2)
	value = append(value, "beijing", "shanghai")
	sliceMap["city"] = value
	for key, val := range sliceMap {
		fmt.Printf("init index: %v value: %v\n", key, val)
	}

	actions := make(map[string]string,3)
	actions["1"] = "test1"
	actions["2"] = "test2"
	actions["3"] = "test3"

	// if 语句可以持初始化语句
	// tip 感觉这里可以用来判断某个元素是否存在
	if action, ok := actions["4"]; ok {
		fmt.Println("===action===>", action, ok)
	} else {
		fmt.Println("===error action===>", action, ok)
	}

	// tip
	//  因为 interface{} 即空接口，它可以存储任意类型，并且任意类型都实现了空接口
	// 	所以我们这样来定义这个字典
	a :=map[string]interface{}{
		"name":"lynn",
		"age":100,
	}
	fmt.Println("a 的值就是===>", a)
}