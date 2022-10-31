package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Hello string

type Args struct {
	Numero int
}

func (t *Hello) SayHello(args *Args, frase *string) (err error) {
	log.Println(args.Numero)
	*frase = "hello world"
	return nil
}
func main() {
	hello := new(Hello)
	rpc.Register(hello)
	rpc.HandleHTTP()

	port := ":9090"
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("error ao escutar a porta")
	}
	log.Printf("eu to te ouvindo arrombado")
	http.Serve(listener, nil)
}
