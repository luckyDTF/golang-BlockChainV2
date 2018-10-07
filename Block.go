package main

import (
	"time"
)

type Block struct {
	//版本
	Verison int64
	//前区块的哈希值
	PrevBlockHash []byte
	//当前区块的哈希值,为了简化代码
	Hash []byte
	//梅克尔根
	MerkelRoot []byte
	//时间戳
	TimeStamp int64
	//难度值
	Bits int64
	//随机值
	Nonce int64

	//交易信息
	Data []byte
}

//func (block *Block) SetHash() {
//	temp := [][]byte{
//		IntToByte(block.Verison),
//		block.PrevBlockHash,
//		block.MerkelRoot,
//		IntToByte(block.TimeStamp),
//		IntToByte(TARGET_BITS),
//		IntToByte(0),
//		block.Data}
//	data := bytes.Join(temp, []byte{})
//	hash := sha256.Sum256(data)
//	block.Hash = hash[:]
//}

func NewBlock(data string, pre []byte) *Block {
	block := Block{
		Verison:       1,
		PrevBlockHash: pre,
		//Hash TODO
		MerkelRoot: []byte{},
		TimeStamp:  time.Now().Unix(),
		Bits:       TARGET_BITS,
		Nonce:      0,
		Data:       []byte(data)}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
