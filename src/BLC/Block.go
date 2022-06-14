package BLC

import (
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Hash          []byte //현재 블록의 해쉬배열
	PrevBlockHash []byte //이전 해쉬값
	Timestamp     int64  //만들어진 시각
	Nonce         int    //난수
	Data          []byte //<---- 나중에 확장하기 MT/MR
}

func (block *Block) setHash() {

	timestamp := strconv.FormatInt(block.Timestamp, 10)
	timeBytes := []byte(timestamp)

	//현재 블록의 해시값 = timeBytes + PrevBlockHash + Data 를 합친 값
	var blockBytes []byte
	blockBytes = append(timeBytes, block.PrevBlockHash...)
	blockBytes = append(blockBytes, block.Data...)
	// 		↳--------------↴
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]

}

func NewBlock(data string, prevBlockHash []byte) *Block {

	block := &Block{}

	//만들어진 시각 , 이전 블록의 해시값 , 블록의 데이터
	block.Timestamp = time.Now().Unix()
	block.PrevBlockHash = prevBlockHash
	block.Data = []byte(data)

	//sethash
	block.setHash()

	pow := newProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
