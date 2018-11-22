package block

import (
	"time"

	blk "simpleblockchain/define/block"
	pow "simpleblockchain/proofofwork"
)

// NewBlock 用于生成新块
// Data : 数据
// PrevBlockHash : 前一个区块Hash
func NewBlock(data string, prevBlockHash []byte) *blk.Block {
	block := new(blk.Block)
	block.Timestamp = time.Now().UnixNano()
	block.PrevBlockHash = prevBlockHash
	block.Data = []byte(data)
	// block.Hash = []byte{}

	// block.SetHash()
	powInfo := pow.NewProofOfWork(block)
	nonce, hash := powInfo.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock : 生成创世块
func NewGenesisBlock() *blk.Block {
	return NewBlock("Genesis Block", []byte{})
}
