package main

import (
	f "fmt"
	"net"
	"net/rpc"
)

type Calc int //RPC 서버에 등록하기 위한 임의의 타입정의
type Args struct {
	A, B int
}
type Reply struct {
	C int
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}
func main() {
	rpc.Register(new(Calc))
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
