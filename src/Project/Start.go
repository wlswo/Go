package Project

import (
	"encoding/json"
	f "fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func (txs *Txs) TxRun(data []byte) {
	txs.AddTx(data)

	for i, tx := range txs.Txs {
		tx.TPrint(i)
	}
}

func (bc *Blockchain) Run(data []byte) {

	//블록 생성 , 블록 추가
	bc.AddBlock(data)

	for i, bc := range bc.Blocks {
		bc.BPrint(i)
	}

	//가변값을 가진 시드를 주어야 math/rand 값이 바뀜
	rand.Seed(time.Now().UnixNano())

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

}

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
