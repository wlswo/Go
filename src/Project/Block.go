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
	Data          Tx     `json:"Data"`          //<---- 나중에 확장하기 MT/MR
	Bits          int    `json:"Bits"`          //Targetbits
	Pow           []byte `json:"Pow"`           //Hash from Pow
	Height        int    `json:"Height"`        //Block Height
}

type Tx struct {
	TxID      []byte //sha256(Data + TimeStamp + Nonce)
	Data      []byte //Do what
	Nonce     int    //Random Num that have Ordering
	TimeStamp int64  //Do job Time
	Sign      []byte //Sign Only Master
}

func (bc *Block) BPrint(i int) {
	f.Println("-------------------------------", i, "번째 블록 데이터 ---------------------------------")
	f.Printf("Prev Hash    	 : %x\n", bc.PrevBlockHash)
	f.Printf("Data             : %s\n", bc.Data)
	f.Printf("hash 		 : %x\n", bc.Hash)
	f.Println("TimeStamp 	 :", time.Unix(bc.Timestamp, 0))
	f.Printf("Nonce 		 : %d\n", bc.Nonce)
	f.Printf("bits		 : %d\n", bc.Bits)
	f.Printf("Height 	 	 : %d\n", bc.Height)
	f.Println("-------------------------------------------------------------------------------------\n")
}

func (block *Block) setTx() {
	timestamp := strconv.FormatInt(block.Timestamp, 10)
	timeBytes := []byte(timestamp)
	Nonce := strconv.FormatInt(block.Data.Nonce, 64)
	NonceBytes := []byte(Nonce)

	var blockBytes []byte
	blockBytes = append(timeBytes, block.Data.Data...)
	blockBytes = append(blockBytes, NonceBytes...)
	blockBytes = append(blockBytes, block.Data.Data...)
	// 		↳--------------↴
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
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
	block.setTx()

	block.Bits = rand.Intn(10)
	pow := newProofOfWork(block)
	nonce, hash := pow.Run()
	block.Pow = hash[:]
	block.Nonce = nonce
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
	return bytes.Equal(b.Data, e)
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
