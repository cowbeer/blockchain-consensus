package pos

import (
	"consensus"
	"math/rand"
	"time"
)

func GenerateNewBlockWithPos(prevBlock consensus.Block, data string, difficulty int) consensus.Block {
	timestamp := time.Now().Unix()
	newBlock := consensus.Block{
		Height:     prevBlock.Height + 1,
		Timestamp:  timestamp,
		PrevHash:   prevBlock.Hash,
		Data:       data,
		Nonce:      0,
		Difficulty: difficulty,
		Hash:       "",
	}

	// select a random validator
	rand.New(rand.NewSource(time.Now().UnixNano()))
	validatorIndex := rand.Intn(len(consensus.Validators))
	validator := consensus.Validators[validatorIndex]
	newBlock.Validator = validator.Address

	newBlock.Hash = consensus.CalculateHash(newBlock)

	return newBlock
}
