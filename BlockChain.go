package main

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	block := NewGenesisBlock()
	return &BlockChain{blocks: []*Block{block}}
}

func (bc *BlockChain) AddBlock(data string) {
	pre := bc.blocks[len(bc.blocks)-1].Hash
	block := NewBlock(data, pre)
	bc.blocks = append(bc.blocks, block)
}
