package main

import "fmt"

func main() {
	sum :=0
	for i := 0; i < 10; i++ {
		if i==sum {
			fmt.Println("======>", sum, i)
			break
		}
		fmt.Println("这里不会执行")
	}
}