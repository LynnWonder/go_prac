package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 本节讲解 Unmarshal

type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class map[string]interface{} `json:"class"`
	// tip
	//  如果我们想让 Class 这种"复合数据"（可能会进行二次甚至更多次 json 解析的）
	//  必须明确类型，否则会被转化成 map[string]interface{}
	//  当然这里也可以用 Class Class1 即普通 struct 类型
	//Class *Class1 `json:"class"`
	Test  interface{}
}


type Class1 struct {
	Name string
	Grade int
}

func main() {
	//json字符中的"引号，需用\进行转义，否则编译出错
	//json字符串沿用上面的结果，但对key进行了大小的修改，并添加了sex数据
	data:="{\"name\":\"张三\",\"Age\":18,\"high\":true,\"sex\":\"男\",\"class\":{\"naME\":\"1班\",\"GradE\":3}}"
	str:=[]byte(data)

	//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//tip
	// Unmarshal 第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象StuRead{}
	//2.可以直接stu:=new(StuRead),此时的stu自身就是指针
	stu:=StuRead{}
	err:=json.Unmarshal(str,&stu)

	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err!=nil{
		fmt.Println(err)
	}

	// tip 关于 Unmarshal
	//  接收的时候，json 串中的 key 的匹配原则是：
	// 	 和标签一样的肯定会匹配
	//   没有 json 标签，那么就从上往下查找变量名或者忽略大小写之后和 key 一样的变量
	// 不可导出的变量无法解析，解析后其值为 nil
	// 有匹配不了的项也是 nil 如本例中的 Test
	// 此外：由于没有指定变量Class的具体类型，json自动将value为复合结构的数据解析为map[string]interface{}类型的项。


	// 打印出来的值：{张三 18 true <nil> map[GradE:3 naME:1班] <nil>}
	fmt.Println(stu)
	fmt.Println(stu.Class,stu.Class["naME"])

	var data0 = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data0, &result); err != nil {
		log.Fatalln(err)
	}

	// 在 encode/decode JSON 数据时，Go 默认会将数值当做 float64
	fmt.Printf("%T\n", result["status"])    // float64
	var status = result["status"]
	fmt.Println("Status value: ", status)
}