package main

import (
	"log"
	"net"
	"net/rpc"
	"runtime"
	"gorpc/context"
)



func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	client := context.NewRPC()
	rpc.Register(client)

	listener, e := net.Listen("tcp", ":9876")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Println("RPC server is listening on port 9876")
	rpc.Accept(listener)



}
