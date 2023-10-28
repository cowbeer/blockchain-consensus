package pow

import (
	"consensus"
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	difficulty := 3
	powBlockchain := []consensus.Block{consensus.CreateGenesisBlockForPow(difficulty)}

	newBlockData := "Block Data"
	newPowBlock := GenerateNewBlockWithPoW(powBlockchain[len(powBlockchain)-1], newBlockData, difficulty)
	powBlockchain = append(powBlockchain, newPowBlock)

	for i := 0; i < 5; i++ {
		prevBlock := powBlockchain[len(powBlockchain)-1]
		newPowBlock = GenerateNewBlockWithPoW(prevBlock, newBlockData, difficulty)
		powBlockchain = append(powBlockchain, newPowBlock)
	}

	fmt.Println("Proof of Work Blockchain:")
	for _, block := range powBlockchain {
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}
