package main

import "fmt"

type Cat struct {
	Name    string
	Age     uint8
	Address  string
}
type Animal interface {
	Grow()
	Move(string) string
}

func (cat *Cat) Grow() {
	cat.Age++
}

func (cat *Cat) Move(add string) string{
	temp :=cat.Address
	cat.Address = add
	return temp
}


func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)
}