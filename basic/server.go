package main

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(ArithService), "")

	s.RegisterBeforeFunc(func(ri *rpc.RequestInfo) {
		fmt.Println(ri)
	})

	s.RegisterAfterFunc(func(ri *rpc.RequestInfo) {
		fmt.Println(ri)
	})

	http.Handle("/rpc", s)

	http.ListenAndServe(":1234", nil)
}

type Args struct {
	A, B int
}

type ArithService struct{}

func (a *ArithService) Add(req *http.Request, args *Args, reply *int) error {
	*reply = args.A + args.B
	fmt.Println("calculating")
	return nil
}
