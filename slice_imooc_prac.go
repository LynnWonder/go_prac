package main
import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// tip 变量的批量声明
	var (
		a string
		b int
	)
	a = "test"
	b = 1
	fmt.Printf("print the val %v, %v", a, b)
	const (
		n1 = 100
		n2
	)
	fmt.Println()
	fmt.Printf("print the val %v, %v", n1, n2)



	fmt.Println("")
	// 获取的是索引 [4,6) 的数据，同时最大索引是 8
	slice5 := numbers4[4:6:8]
	// 子切片 [start, end)  s[low:high:max] len=high-low cap=max-low
	fmt.Print("====>", slice5)
	for i:=range slice5{
		slice5[i]+=5
	}
	// range() 函数可以用来遍历数组、切片和字典
	for i,v:=range slice5{
		fmt.Printf("键：%v：值：%v\n", i,v)
	}
	// Tip 由于 slice 是引用类型因此导致了改变了原先的数组
	for i,v:=range numbers4{
		fmt.Printf("改变了键：%v：值：%v\n", i,v)
	}
	length := 2
	capacity := 4
	fmt.Printf("%v, %v, %v\n", slice5,length == len(slice5), capacity == cap(slice5))
	// key: 因为容量拉大，而切片实际上是引用类型，此时 slice5 是 10,11,7,8
	slice5 = slice5[:cap(slice5)]
	fmt.Print("====>", slice5)
	slice5 = append(slice5, 11, 12, 13)
	length = 7
	fmt.Printf("%v, %v\n", length == len(slice5), slice5)
	// key: 此时 slice5 = [10,11,7,8,11,12,13]
	slice6 := []int{0, 0, 0}
	// func copy(dst, src []Type) int Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
	copyLen := copy(slice5, slice6)
	e2 := 0
	e3 := 8
	e4 := 11
	fmt.Printf("%v, %v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4], copyLen)
}