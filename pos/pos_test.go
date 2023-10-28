package pos

import (
	"consensus"
	"fmt"
	"testing"
)

func TestPos(t *testing.T) {
	difficulty := 3
	posBlockchain := []consensus.Block{consensus.CreateGenesisBlockForPos(difficulty)}
	newBlockData := "Block Data"

	newPosBlock := GenerateNewBlockWithPos(posBlockchain[len(posBlockchain)-1], newBlockData, difficulty)
	posBlockchain = append(posBlockchain, newPosBlock)

	for i := 0; i < 5; i++ {
		prevBlock := posBlockchain[len(posBlockchain)-1]
		newPosBlock = GenerateNewBlockWithPos(prevBlock, newBlockData, difficulty)
		posBlockchain = append(posBlockchain, newPosBlock)
	}

	fmt.Println("Proof of Stake Blockchain:")
	for _, block := range posBlockchain {
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Validator: %s\n", block.Validator)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}
