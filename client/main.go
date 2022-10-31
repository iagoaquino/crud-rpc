package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	Numero int
}

func main() {
	host := "localhost:9090"
	var hello string
	cliente, err := rpc.DialHTTP("tcp", host)

	if err != nil {
		log.Fatal("fudeo tudo menor")
	}

	err = cliente.Call("Hello.SayHello", args, &hello)
	if err != nil {
		log.Fatal("agora fudeo tudo no final", err)
	}
	log.Printf("%s\n", hello)
}
