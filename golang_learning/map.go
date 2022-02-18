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


	// map 是否有 JSON.stringify 这样的函数


	// 操作 map 字典
	delete(m2, "Sam")
	fmt.Println("m2", m2)
}