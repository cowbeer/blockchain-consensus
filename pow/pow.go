package pow

import (
	"consensus"
	"math"
	"math/big"
	"time"
)

func GenerateNewBlockWithPoW(prevBlock consensus.Block, data string, difficulty int) consensus.Block {
	timestamp := time.Now().Unix()
	newBlock := consensus.Block{
		Height:     prevBlock.Height + 1,
		Timestamp:  timestamp,
		PrevHash:   prevBlock.Hash,
		Data:       data,
		Nonce:      0,
		Difficulty: difficulty,
	}

	var nonce int64
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	for nonce < math.MaxInt64 {
		newBlock.Nonce = nonce
		hash := consensus.CalculateHash(newBlock)
		hashInt := new(big.Int)
		hashInt.SetString(hash, 16)
		if hashInt.Cmp(target) == -1 { // hashInt < target
			newBlock.Hash = hash
			break
		} else {
			nonce++
		}
	}

	return newBlock
}
