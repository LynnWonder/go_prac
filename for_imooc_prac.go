package main

import (
	"fmt"
)

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for i,v := range map1 {
		fmt.Printf("%d: %v\n", i, v)
	}
}