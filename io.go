package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 写文件和读取文件
	arr :=[]byte{'b', 'n','t'}
	ioutil.WriteFile("state.yaml", arr, 0666)
	inbytes, err := ioutil.ReadFile("state.yaml")
	fmt.Println(inbytes)
	if err != nil {
		fmt.Println(err)
	}
}


