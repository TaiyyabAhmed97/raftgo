package main

import (
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		panic("error in connecting to server")
	}
	var reply string
	var args = "hello from client1"
	err = client.Call("Node.PrintNode", args, &reply)
	if err != nil {
		panic("erorr in calling rpc server methpd")
	}
}
