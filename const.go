package main

import "fmt"

func getNum() int {
	return 2
}
func main() {
	// TIP golang 常量：必须是在编译的时候就能确定的值，可以通过计算得到
	const c1 = 2 / 3
	// TIP 以下将会报错
	//const c2 = getNum()
	// TIP 格式化说明符：
	//  %d 格式化整数  （%x 和 %X 用于格式化 16 进制表示的数字）
	//  %g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法）
	//  %0nd 用于规定输出长度为n的整数，其中开头的数字 0 是必须的
	//  %c 用于表示字符
	//  %s 表示字符串
	//  %v 表示以变量自己的类型默认输出格式输出
	fmt.Printf("c1===>%d", c1)
}
