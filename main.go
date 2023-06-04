package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Node string

func (n *Node) PrintNode(nodeid string, reply *string) error {
	fmt.Printf("nodeid args: %v and reply arg: %v\n", nodeid, *reply)
	*reply = nodeid
	fmt.Printf("Update reply arg: %v", *reply)
	return nil
}

func main() {
	node := new(Node)
	rpc.Register(node)

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		// do something with error
		panic("womething is worng with the netowkr")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// domsething
		}

		go rpc.ServeConn(conn)
	}
}
