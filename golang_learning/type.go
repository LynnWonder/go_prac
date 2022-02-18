package main

import (
	"fmt"
	"reflect"
)
/**
 * 关于类型转换 https://juejin.cn/post/6844904113914789902
 * http://www.randyfield.cn/post/2021-05-26-go-type-conversion/
 */
func  main()  {
	// ① golang 中空接口可以存储任意类型的值，因此可以使用空接口来进行数据转换
	var a interface{} = 257
	var b interface{} = "zz"

	s,ok := a.(int64)
	z := b.(string)
	// ② unsafe 强制类型转换，暂时不学习
	//var f float64
	//bits = *(*uint64)(unsafe.Pointer(&f))
	fmt.Println(s, ok)
	fmt.Println(z)
	fmt.Println(s, reflect.TypeOf(s))
	justifyType(nil)
	justifyType("I'm Garfield")
	justifyType(44)
	justifyType(int64(516165161616))
	justifyType(true)
	justifyType(int32(1105020))
}



// ① switch:变量 x 断言成了 type 类型，type 类型具体值就是 switch case 的值，如果 x 成功断言成了某个 case 类型，就可以执行那个 case，此时 i := x.(type) 返回的 i 就是那个类型的变量了，可以直接当作 case 类型使用。
func justifyType(x interface{}) {
	switch i := x.(type) {
	case int64:
		fmt.Printf("x is a int64, is %v\n", i)
	case string:
		fmt.Printf("x is a string，value is %v\n", i)
	case int:
		fmt.Printf("x is a int is %v\n", i)
	case bool, int32:
		fmt.Printf("x is a bool or int32 is %v\n", i)
	case nil:
		fmt.Printf("x is a interface{} is %v\n", i)
	default:
		fmt.Println("unsupport type！")
	}
}