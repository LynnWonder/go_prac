package main

import (
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	origin := uuid.NewV4()
	u := origin.Bytes()
	fmt.Println(origin)
	buf := make([]byte, 16)
	hex.Encode(buf[0:12], u[0:6])
	hex.Encode(buf[12:], u[8:10])
	fmt.Println(string(buf))
	fmt.Println(len(string(buf)))
}
