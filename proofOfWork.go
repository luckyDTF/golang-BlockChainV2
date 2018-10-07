package main

import (
	"math/big"
	"bytes"
	"math"
	"crypto/sha256"
	"fmt"
)

type ProofOfWork struct {
	block *Block
	//目标值 比较hash值
	target *big.Int
}

const TARGET_BITS = 24

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TARGET_BITS))
	pow := ProofOfWork{block: block, target: target}
	return &pow
}

func (pow *ProofOfWork) PrepareData(nonce int64) []byte {
	block := pow.block
	temp := [][]byte{
		IntToByte(block.Verison),
		block.PrevBlockHash,
		block.MerkelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(TARGET_BITS),
		IntToByte(nonce),
		block.Data}
	data := bytes.Join(temp, []byte{})
	return data
}

//随机hash进行比较-挖矿
func (pow *ProofOfWork) Run() (int64, []byte) {
	//1.拼装数据
	//2.哈希值转成big.Int类型
	var hash [32]byte
	var hashInt big.Int
	var nonce int64 = 0

	fmt.Println("Begin Mining...")
	fmt.Printf("target hash : %x\n", pow.target.Bytes())

	for ; nonce < math.MaxInt64; {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("found hash ,hash :%x, nonce : %d\n", hash, nonce)
			break
		} else {
			//fmt.Printf("not found nonce ,current nonce :%d, hash : %x\n", nonce, hash)
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}
