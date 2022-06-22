package Project

import (
	"encoding/json"
	f "fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func (bc *Blockchain) Run(str string) {
	// 블록체인 생성
	// 리턴값은 BlockChain 구조체의 주소를 반환
	//bc := NewBlockchain()

	bc.AddBlock(str)

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
				f.Printf("%s\n", v.Data.Data)
				break
			}
		}
		//block
		inHash = block.PrevBlockHash

		if block.IsGenBlock() {
			f.Println("Completed block traversal but not found\n")
			break
		}
	}

	//Copy Origin Block struct then Paste Format of Json
	block_doc, _ := json.MarshalIndent(bc.Blocks, "", " ") //BlockChain Blocks
	err := ioutil.WriteFile("blockFile.json", block_doc, os.FileMode(0644))

	if err != nil {
		f.Println(err)
		return
	}

	b, err := ioutil.ReadFile("blockFile.json")

	if err != nil {
		f.Println(err)
		return
	}

	var Copychain []*Copy         // save the bytes slice for read
	json.Unmarshal(b, &Copychain) // read Json file

	//Compare To Data
	//f.Printf("%t", CompareBlock(bc.Blocks, b))

}

/*
func CompareBlock(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
*/

//copy Struct
type Copy struct {
	Hash          []byte `json:"Hash"`          //Current Hash
	PrevBlockHash []byte `json:"PrevBlockHash"` //Previous Hash
	Timestamp     int64  `json:"Timestamp"`     //Block was create To Time
	Nonce         int    `json:"Nonce"`         //Random Num that have Ordering
	Data          struct {
		TxID      []byte `json:"TxID"`      //sha256(Data + TimeStamp + Nonce)
		Data      []byte `json:"Data"`      //Do what
		Nonce     int    `json:"Nonce"`     //Random Num that have Ordering
		TimeStamp int64  `json:"TimeStamp"` //Do job Time
		Sign      []byte `json:"Sign"`      //Sign Only Master
	}
	Bits   int    `json:"Bits"`   //Targetbits
	Pow    []byte `json:"Pow"`    //Hash from Pow
	Height int    `json:"Height"` //Block Height
}
