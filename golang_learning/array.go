package main

import (
	"fmt"
	"reflect"
)

func main()  {
	// 声明数组
	//var arr [6]int
	// 声明时初始化
	var arr = [5]int{1,2,3,4,5}
	arr1 :=[2]int{1,2}
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(arr1)


	// 修改数组
	for i :=0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)


	// 切片是数组的抽象。 切片使用数组作为底层结构。切片包含三个组件：容量，长度和指向底层数组的指针,切片可以随时进行扩展
	// 切片就是更加灵活的数组，也可以理解为就是动态数组
	// 长度为 3 容量为 5 的切片
	slice1 :=make([]float32, 3, 5)
	// 此时值为 [0,0,0]
	fmt.Println(slice1)


	// 数组拼接，数组切割（获取数组的一部分）
	s1 := arr[2:5:5]
	fmt.Println(s1, len(s1), cap(s1), reflect.TypeOf(s1), reflect.TypeOf(arr))
	for i:=0;i<len(s1);i++ {
		s1[i] -=100
	}
	// 由此可见更改了数组切割下来的数据后原数组也会发生变化
	fmt.Println(s1, len(s1), cap(s1), reflect.TypeOf(s1), arr, reflect.TypeOf(arr))
	// 切片拼接
	var a = []int{1,2,3}
	a = append(a, []int{4}...)
	a = append([]int{0}, a...)
	b := a[:len(a)-1]
	for i :=range b {
		b[i] +=10
	}
	// 可以发现给 b 赋值 a 切掉最后一个元素后的数组，a 并没有变
	// 但是修改 b 之后 a 还是发生了变化
	fmt.Println("切片截取",a, b)
	// 可以使用 copy 来做切片的各种变化



	// 数组遍历
	// 使用 for range 迭代的性能更好一些，这种迭代保证不会出现数据越界的情形，因此每轮迭代也就省去了对下标是否越界的判断
	for i :=range arr {
		// 整型输出
		fmt.Printf("arr[%d]: %d\n", i, arr[i])
	}

	for i,v :=range arr {
		fmt.Printf("arr[%d]: %d\n", i, v)
	}

	for i:=0;i<len(arr);i++ {
		fmt.Printf("arr[%d]: %d\n", i, arr[i])
	}

}