package Project

import (
	"encoding/json"
	f "fmt"
	"io/ioutil"
	"os"
)

var cnt = 0

func (txs *Txs) CreateTx(data []byte) {
	//트랜잭션 생성
	txs.AddTx(data)

	tx := txs.Txs[len(txs.Txs)-1]
	tx.TPrint(cnt)
	cnt++

	Tx_doc, _ := json.MarshalIndent(txs.Txs, "", " ") //BlockChain Blocks
	err := ioutil.WriteFile("TxFile.json", Tx_doc, os.FileMode(0644))

	if err != nil {
		f.Println(err)
		return
	}
}

func (bc *Blockchain) CreateBc(data []byte) {

	//블록 생성 , 블록 추가
	bc.AddBlock(data)

	//Copy Origin Block struct then Paste Format of Json
	block_doc, _ := json.MarshalIndent(bc.Blocks, "", " ") //BlockChain Blocks
	err := ioutil.WriteFile("BlockFile.json", block_doc, os.FileMode(0644))

	if err != nil {
		f.Println(err)
		return
	}

	b, err := ioutil.ReadFile("BlockFile.json")

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
