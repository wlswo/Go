package main

import (
	"BLC"
	//"bytes"
	f "fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// 블록체인 생성
	// 리턴값은 BlockChain 구조체의 주소를 반환
	bc := BLC.NewBlockchain()

	for i := 0; i < 5; i++ {
		idx := strconv.Itoa(rand.Intn(1000))
		bc.AddBlock(idx)
	}

	for i, bc := range bc.Blocks {
		bc.BPrint(i)
	}

	//가변값을 가진 시드를 주어야 math/rand 값이 바뀜
	rand.Seed(time.Now().UnixNano())
	request := []byte(strconv.Itoa(rand.Intn(1000)))

	//최상의 블럭의 PrevBlockHash
	inHash := bc.Blocks[len(bc.Blocks)-1].Hash

	for _, v := range bc.Blocks {

		block := bc.FindBlock(inHash)

		if block != nil {
			if v.EqualData(request) {
				f.Println("found")
				f.Printf("%s\n", v.Data)
				break
			}
		}
		//block
		inHash = block.PrevBlockHash

		if block.IsGenBlock() {
			f.Println("Completed block traversal but not found")
			break
		}

	}

}
