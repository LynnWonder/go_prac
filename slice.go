package main

import "fmt"

// tip 声明切片的几种方式
func main() {
	// 1.声明切片
	// tip 数组定义：
	//  var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。
	//  一旦定义，长度不能变。
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 切片追加值
	s1 = append(s1, 2)
	fmt.Println("=====>s1",s1)

	// 2.:=
	s2 := []int{}

	// 3.make()，make([]int, len, cap)，若省略 cap 则 cap=len
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	// 数组也可以这样赋值，只不过会添上长度或者 ...
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)


	s := []int{0, 1, 2, 3}
	p := &s[2] // *int, tip 获取底层数组元素指针。
	*p += 100 // 对指针所指的位置的值增加 100
	// TODO 了解指针
	fmt.Println("===s 变化后===",s)


	// append 操作切片，向 slice 尾部添加数据，返回新的 slice 对象
	// tip 当 append 超出了原 slice.cap 限制，就会重新分配底层数组，即使原数组没有被填满
	var a = []int {1,2,3}
	var b = []int{4,5,6}
	c := append(a,b...)
	fmt.Printf("slice c: %v\n", c)
	d :=append(c,7)
	fmt.Printf("slice c,d: %v, %v\n", c, d)


	// 切片拷贝
	data := [...]int{0,1,2,3,4,5,6,7,8,9}
	ss1 := data[8:] // [8,9]
	ss2 := data[:5] // [0,1,2,3,4]
	copy(ss2,ss1)
	fmt.Println("====>ss1", ss1)
	fmt.Println("====>ss2", ss2)


	// slice 遍历
	for index,val := range ss1 {
		fmt.Println("====>ss1 ele",index ,val)
	}

	// slice resize
	var aa = []int{1, 3, 4, 5}
	fmt.Printf("slice aa : %v , len(aa) : %v\n", aa, len(aa))
	bb := aa[1:2]
	fmt.Printf("slice bb : %v , len(bb) : %v\n", bb, len(bb))
	cc := bb[0:3]
	fmt.Printf("slice cc : %v , len(cc) : %v\n", cc, len(cc))

	// 字符串和切片 string 底层是一个 byte 的数组
	strr := "hello world"
	// 含有中文字符的字符串则使用 []rune(strr)
	ss := []byte(strr)
	ss[6] = 'Z'
	ss = append(ss, '!')
	fmt.Println("=====>变化后的字符串", string(ss))


}