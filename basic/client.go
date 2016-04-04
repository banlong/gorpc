package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/json"
	"net/http"
)

func main() {
	args := Args{4, 13}
	msg, _ := json.EncodeClientRequest("ArithService.Add", args)
	resp, _ := http.Post("http://localhost:1234/rpc", "application/json",
		bytes.NewReader(msg))

	var result int
	json.DecodeClientResponse(resp.Body, &result)

	fmt.Println(result)
}

type Args struct {
	A, B int
}
