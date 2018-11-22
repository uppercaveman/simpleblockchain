package proofofwork

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"

	blk "simpleblockchain/define/block"
	"simpleblockchain/modules/util"
)

// 难度值，这里表示哈希的前24位必须是0,数字越大难道越大
const targetBits = 1

const maxNonce = math.MaxInt64

// ProofOfWork : 每个块的工作量都必须要证明，所有有个指向 Block 的指针
type ProofOfWork struct {
	Block  *blk.Block // 区块
	Target *big.Int   // target是目标，我们最终要找的哈希必须要小于目标
}

// NewProofOfWork : 创建矿工
func NewProofOfWork(b *blk.Block) *ProofOfWork {
	// target 等于 1 左移 256 - targetBits 位
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// 工作量证明用到的数据有: PrevBlockHash, Data, Timestamp, targetBits, nonce
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	timestamp, err := util.Int64ToByte(pow.Block.Timestamp)
	if err != nil {
		panic(err)
	}
	targetBits, err := util.Int64ToByte(int64(targetBits))
	if err != nil {
		panic(err)
	}
	nonceByte, err := util.Int64ToByte(int64(nonce))
	if err != nil {
		panic(err)
	}

	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			timestamp,
			targetBits,
			nonceByte,
		},
		[]byte{},
	)

	return data
}

// Run : 工作量证明的核心就是寻找有效哈希
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	// log.Info("Mining the block containing : %s", pow.Block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			// log.Info("ProofOfWork hash: %x", hash)
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// Validate : 验证工作量，只要哈希 小于 目标 就是有效工作量
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.Target) == -1

	return isValid
}
