package Project

import (
	"bytes"
	"crypto/sha256"
	f "fmt"
	"math/rand"
	"strconv"
	"time"
)

type Block struct {
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

type Tx struct {
}

func (bc *Block) BPrint(i int) {
	f.Println("-------------------------------", i, "번째 블록 데이터 ---------------------------------")
	f.Printf("Prev Hash    	 : %x\n", bc.PrevBlockHash)
	f.Println("TimeStamp 	 :", time.Unix(bc.Timestamp, 0))
	f.Printf("Nonce 		 : %d\n", bc.Nonce)
	f.Printf("bits		 : %d\n", bc.Bits)
	f.Printf("hash 		 : %x\n", bc.Hash)
	f.Printf("Height 	 	 : %d\n", bc.Height)
	f.Println("Transaction  { 	 ")
	f.Printf("	TxID	:  %x\n", bc.Data.TxID)
	f.Printf("	Data	:  %s\n", bc.Data.Data)
	f.Printf("	Nonce	:  %d\n", bc.Data.Nonce)
	f.Println("	TimeStamp : ", time.Unix(bc.Data.TimeStamp, 0))
	f.Printf("	Sign	:  %x\n", bc.Data.Sign)
	f.Println("} ")
	f.Println("-------------------------------------------------------------------------------------\n")
}

func (block *Block) setTx(data string) {
	block.Data.Data = []byte(data)

	timestamp := strconv.FormatInt(block.Timestamp, 10)
	timeBytes := []byte(timestamp)
	block.Data.TimeStamp = time.Now().Unix()

	Nonce := strconv.Itoa(block.Data.Nonce)
	NonceBytes := []byte(Nonce)

	var blockBytes []byte
	blockBytes = append(timeBytes, block.Data.Data...)
	blockBytes = append(blockBytes, NonceBytes...)
	blockBytes = append(blockBytes, block.Data.Data...)
	// 		↳--------------↴
	hash := sha256.Sum256(blockBytes)
	block.Data.TxID = hash[:]
	block.Data.Sign = hash[:]
}

func (block *Block) setHash() {

	timestamp := strconv.FormatInt(block.Timestamp, 10)
	timeBytes := []byte(timestamp)

	//현재 블록의 해시값 = timeBytes + PrevBlockHash + Data 를 합친 값
	var blockBytes []byte
	blockBytes = append(timeBytes, block.PrevBlockHash...)
	blockBytes = append(blockBytes, block.Pow...)
	blockBytes = append(blockBytes, block.Data.Data...)
	// 		↳--------------↴
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]

}

var Height = 0

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{}

	//만들어진 시각 , 이전 블록의 해시값 , 블록의 데이터
	block.Timestamp = time.Now().Unix()
	block.PrevBlockHash = prevBlockHash

	//sethash
	block.setHash()

	//setTx
	block.setTx(data)

	block.Bits = rand.Intn(10)
	pow := newProofOfWork(block)
	nonce, hash := pow.Run()
	block.Pow = hash[:]
	block.Nonce = nonce
	block.Data.Nonce = nonce
	block.Height = Height
	Height++
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

//해시값의 체이닝 여부
func (b *Block) EqualHash(e []byte) bool {
	return bytes.Equal(b.Hash, e)
}

//Data의 정합성 여부
func (b *Block) EqualData(e []byte) bool {
	return bytes.Equal(b.Data.Data, e)
}

//블록찾기
func (bc *Blockchain) FindBlock(id []byte) *Block {
	for _, v := range bc.Blocks {
		if v.EqualHash(id) {
			return v
		}
	}
	return nil
}

func (b *Block) IsGenBlock() bool {
	return bytes.Equal(b.PrevBlockHash, make([]byte, len(b.PrevBlockHash)))

}
