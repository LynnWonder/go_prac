package main
import "fmt"

/**
%d：表示bai把数据按十进制整型输出。

%o：表示把数据按八进制整型输出。

%x：表示把数据按十六进制整型输出。

%u：表示把数据参数按无符号整型输出。

%f：显示小数表示的普通浮点数。
 */

// go 数组
func main()  {
	// 因为数组的长度是数组类型的一个部分，不同长度或不同类型的数据组成的数组都是不同的类型，因此在Go语言中很少直接使用数组（不同长度的数组因为类型不同无法直接赋值）。
	var a [ 3 ] int // 定义一个长度为3的int类 型数组, 元素全部为0
	var b = [...] int { 1 , 2 , 3 } // 定义一个长度为3的 int类型数组, 元素为 1, 2, 3
	var c = [...] int { 2 : 3 , 1 : 2 } // 定义一个长度为3 的int类型数组, 元素为 0, 2, 3
	var d = [...] int { 1 , 2 , 4 : 5 , 6 } // 定义一个长度为 6的int类型数组, 元素为 1, 2, 0, 0, 5, 6
	fmt.Println("a====>",a)
	fmt.Println("b====>",b)
	fmt.Println("c====>",c)
	fmt.Println("d====>",d)
	// 多维数组表示
	//var arr [5]int     // 一维
	//var arr2 [5][5]int // 二维
	// 或 arr := [5]int{1, 2, 3, 4, 5}

	arr := [5]int{1, 2, 3, 4, 5}
	// for 循环迭代数组的一种方式
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)  // [101 102 103 104 105]





	// 指向数组的指针
	var dd = &d
	// 此时依然可以通过 for range 迭代数组指针指向的数组元素，
	for i,v:=range dd{
		fmt.Printf("dd[%d]: %d\n", i, v)
	}
	for i:=range dd{
		fmt.Println("键",i)
	}
	// 切片是数组的抽象，与数组不同的是无法通过切片类型来确定其值的长度。
	// 切片使用数组作为底层结构。
	// 切片包含三个组件：容量，长度和指向底层数组的指针,切片可以随时进行扩展
	slice1 := make([]float32, 0) // 长度为0的切片
	// key: 设置不同容量但长度不一样的切片
	slice2 := make([]float32, 3, 5) // [0 0 0] 长度为3容量为5的切片
	// cap 函数可以用于计算数组的容量。
	// 不过对于数组类型来说， len 和 cap 函数返回的结果始终 是一样的，都是对应数组类型的长度
	fmt.Printf("slice1 长度：[%d], 容量：[%d]", len(slice1), cap(slice1))   // 3 5
	fmt.Println()
	fmt.Printf("slice2 长度：[%d], 容量：[%d]\n", len(slice2), cap(slice2)) // 3 5
	// 添加元素，切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4) // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(len(slice2), cap(slice2)) // 7 12

	// 子切片 [start, end)
	sub1 := slice2[3:] // [1 2 3 4]
	sub2 := slice2[:3] // [0 0 0]
	//sub3 := slice2[3:4]
	// 合并切片
	combined := append(sub1, sub2...)
	combined1 := append([]int{1,2,3}, []int{1,2,3}...)
	fmt.Printf("combined:%v", combined)
	fmt.Println()
	fmt.Printf("combined1:%v", combined1)

}