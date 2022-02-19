package main

import (
	"example/calc"
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(calc.Add(1,2))
}