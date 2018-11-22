package blockchain

import (
	"simpleblockchain/block"
	d "simpleblockchain/define/block"
)

// Blockchain : Block指针数组
type Blockchain struct {
	Blocks []*d.Block
}

// AddBlock : 添加新区块到区块链
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain : 创建一个带有创世区块的Blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*d.Block{block.NewGenesisBlock()}}
}
