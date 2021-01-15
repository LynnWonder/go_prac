package main

import "fmt"

func main()  {
	var number int
	if number := 4; 100 > number {
		number += 3
		fmt.Println(number)
	} else if 100 < number {
		number -= 2
	} else {
		fmt.Println("OK!")
	}
	fmt.Println(number)
}