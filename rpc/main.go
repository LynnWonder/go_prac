package main

import (
	api "github.com/LynnWonder/gin_prac/rpc/kitex_gen/api/echo"
	"log"
)

func main() {
	svr := api.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
