package main
import "fmt"

// todo 重构代码
// go 数组
func main()  {
	//var arr [5]int     // 一维
	//var arr2 [5][5]int // 二维
	//var arr = [5]int{1, 2, 3, 4, 5}
	// 或 arr := [5]int{1, 2, 3, 4, 5}

	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)  // [101 102 103 104 105]
// 切片是数组的抽象。
//切片使用数组作为底层结构。
//切片包含三个组件：容量，长度和指向底层数组的指针,切片可以随时进行扩展
//	slice1 := make([]float32, 0) // 长度为0的切片
	slice2 := make([]float32, 3, 5) // [0 0 0] 长度为3容量为5的切片
	// cap 函数可以用于计算数组的容量。
	// 不过对于数组类型来说， len 和 cap 函数返回的结果始终 是一样的，都是对应数组类型的长度
	fmt.Println(len(slice2), cap(slice2)) // 3 5
	// 添加元素，切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4) // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(len(slice2), cap(slice2)) // 7 12

	// 子切片 [start, end)
	sub1 := slice2[3:] // [1 2 3 4]
	sub2 := slice2[:3] // [0 0 0]
	sub3 := slice2[1:4] // [0 0 1]
	// 合并切片
	combined := append(sub1, sub2...) // [1, 2, 3, 4, 0, 0, 0]
}