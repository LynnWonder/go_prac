package main

import "fmt"

func main() {
	// 仅声明
	m1 := make(map[string]int)
	// 声明时初始化
	m2 := map[string]string{
	"Sam": "Male",
	"Alice": "Female",
	}
	// 赋值/修改
	m1["Tom"] = 18
	fmt.Println("m1", m2)
	fmt.Println("m2", m2)

	// tip interface{} 空数组可以表示任何类型
	m3 := make(map[string]interface{})
	m3["a"] = "a"
	m3["b"] = 10
	m3["c"] = []int{1,2,3}
	fmt.Println("任意类型====>", m3)

	// 操作 map 字典
	delete(m2, "Sam")
	fmt.Println("m2", m2)
}