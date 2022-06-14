package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//블록 구조
type Block struct {
	Timestamp     int64  //블록이 만들어진 시간
	Data          []byte //블록에 담길 데이터 ( 실제 가치를 지닌 정보 )
	PrevBlockHash []byte //이번 블록의 해쉬값
	Hash          []byte //현재 블록의 해쉬값
}

//블록에 해쉬값을 계산
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//매개변수로는 data 의 문자형 과 이전 블록의 해쉬를 바이트값으로 받아와서
//timestamp 와 현재 블록의 데이터와 현재블록의 해쉬값만 넣어주면 블록은 생성된다.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

/* 블록체인 구현 */
//블록들을 값으로 삼는 배열 구조체 -> 간단한 체인
type Blockchain struct {
	blocks []*Block
}

//블록 추가
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]   //총 블록의 길이 -1
	newBlock := NewBlock(data, prevBlock.Hash) //새로운 블럭의 데이터와 이전 해쉬값
	bc.blocks = append(bc.blocks, newBlock)    //blocks 배열에 저장
}

//제네시스 블록 정의
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//제네시스 블록 배열에 투입
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
