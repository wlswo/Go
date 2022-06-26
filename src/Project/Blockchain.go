package Project

type Blockchain struct {
	Blocks []*Block `json:"Blocks`
}

func (blockchain *Blockchain) AddBlock(TxID []byte) {
	// AddBlock 이 처리되는 로직순서
	// NewBlock() (block 리턴)  -> BlockChain 구조체에 저장

	//이전 블록 take
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	//새로운 블록생성 넘길값 data , prevBlockHash
	newBlock := NewBlock(TxID, prevBlock.Hash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
