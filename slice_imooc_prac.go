package main
import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 获取的是索引 [4,6) 的数据，同时最大索引是 8
	slice5 := numbers4[4:6:8]
	for i:=range slice5{
		slice5[i]+=5
	}
	for i,v:=range slice5{
		fmt.Printf("键：%v：值：%v\n", i,v)
	}
	// key: 由于 slice 是引用类型因此导致了改变了原先的数组
	for i,v:=range numbers4{
		fmt.Printf("改变了键：%v：值：%v\n", i,v)
	}
	length := 2
	capacity := 4
	fmt.Printf("%v, %v, %v\n", slice5,length == len(slice5), capacity == cap(slice5))
	// key: 因为容量拉大，而切片实际上是引用类型，此时是 5,6,7,8
	slice5 = slice5[:cap(slice5)]
	slice5 = append(slice5, 11, 12, 13)
	length = 7
	fmt.Printf("%v\n", length == len(slice5))
	// key: 此时 slice5 = [0,0,0,8,11,12,13]
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6)
	e2 := 0
	e3 := 8
	e4 := 11
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
}