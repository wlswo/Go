package main

import (
	f "fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Reply struct {
	C int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:6000")
	if err != nil {
		f.Println(err)
		return
	}
	defer client.Close()

	args := &Args{1, 2}
	reply := new(Reply)
	err = client.Call("Calc.Sum", args, reply) //Calc.Sum
	if err != nil {
		f.Println(err)
		return
	}
	f.Println(reply.C)

	args.A = 4
	args.B = 9
	sumCall := client.Go("Calc.Sum", args, reply, nil)
	<-sumCall.Done
	f.Println(reply.C)

}
