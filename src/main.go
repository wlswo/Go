package main

import (
	"BLC"
	f "fmt"
)

func main() {
	bc := BLC.NewBlockchain()
	bc.AddBlock("1번째 트랜잭션")
	bc.AddBlock("2번째 트랜잭션")
	for i, bc := range bc.Blocks {
		f.Println("-------------------------------", i, "번째 블록 데이터 --------------------------------")
		f.Printf("이전블록 해시    : %x\n", bc.PrevBlockHash)
		f.Printf("데이터           : %s\n", bc.Data)
		f.Printf("현재 블럭의 hash : %x\n", bc.Hash)
		f.Println("-----------------------------------------------------------------------------------\n")
	}
}
