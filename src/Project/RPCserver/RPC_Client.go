package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/gob"
	f "fmt"
	"net/rpc"
)

type Args struct {
	A int
}
type Wallet struct {
	Timestamp  int64
	Address    string
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func main() {
	gob.Register(elliptic.P256())
	client, err := rpc.Dial("tcp", "localhost:6000")
	if err != nil {
		f.Println(err)
		return
	}
	defer client.Close()

	args := &Args{0}
	reply := new(Wallet)
	err = client.Call("Calc.Get", args, reply) //func 호출

	if err != nil {
		f.Println(err)
		return
	}
	f.Printf("Private Key : %x\n", reply.PrivateKey)
	f.Printf("PublicKey : %x\n", reply.PublicKey)
	f.Printf("Wallet Addr : %s\n", reply.Address)
}
