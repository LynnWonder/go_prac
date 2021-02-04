package main

import (
	"fmt"
	"strings"
)

func main()  {
	args :=strings.Split("waterMark text=this is a test dsfd", " ")
	var other = make([]string, 2)
	if strings.Split(args[len(args)-1], "=")[0] == "size" {
		other[0] = strings.Join(args[1: len(args)-1], " ")
		other[1] = args[len(args)-1]
	}else{
		other[0] = strings.Join(args[1:], " ")
	}
	//c :=strings.Join(args[1: len(args)-1], " ")
	fmt.Printf("c====> %v\n", other)
	fmt.Printf("c====> %v\n", other[0])
	fmt.Printf("c====> %v\n", other[1])
	for i := 0; i < len(other); i++ {
		fmt.Printf("c====> %v\n", other[i])
		xx := strings.Split(other[i], "=")
		fmt.Printf("xx====> %v\n", xx)
	}
	//fmt.Printf("args====> %v\n", c[1])
	//fmt.Printf("args====> %v\n", args)
	//fmt.Printf("args====> %v\n", args[0])
	//fmt.Printf("args====> %v\n", args[1])
	//fmt.Printf("args====> %v\n", args[2])
	//fmt.Printf("args====> %v\n", args[3])
	//fmt.Printf("args====> %v\n", args[4])
	//fmt.Printf("args====> %v\n", args[len(args)-1])
}
