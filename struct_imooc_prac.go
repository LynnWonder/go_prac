package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func (person *Person) Grow() {
	person.Age++
}

func (person *Person) Move(add string) string{
	temp :=person.Address
	person.Address = add
	return temp
}

// expected to print Robert moved from Beijing to San Francisco.

func main() {
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}