package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

// Block : 区块
type Block struct {
	Timestamp     int64  // Timestamp : 时间戳，也就是区块创建的时间
	PrevBlockHash []byte // PrevBlockHash : 前一个块的哈希
	Hash          []byte // Hash : 当前块的哈希
	Data          []byte // Data : 区块实际存储的信息，比特币中也就是交易
	Nonce         int    // Nonce 在对工作量证明进行验证时用到
}

// SetHash : 设置当前块哈希
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// Hash = sha256(PrevBlockHash + Data + Timestamp)
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
