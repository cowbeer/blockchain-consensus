package consensus

import "time"

func CreateGenesisBlockForPow(difficult int) Block {
	timestamp := time.Now().Unix()
	genesisBlock := Block{
		Height:     1,
		Timestamp:  timestamp,
		PrevHash:   "0",
		Data:       "Genesis Block",
		Nonce:      0,
		Difficulty: difficult,
	}

	genesisBlock.Hash = CalculateHash(genesisBlock)
	return genesisBlock
}

func CreateGenesisBlockForPos(difficulty int) Block {
	timestamp := time.Now().Unix()
	genesisBlock := Block{
		Height:     1,
		Timestamp:  timestamp,
		PrevHash:   "0",
		Data:       "Genesis block",
		Nonce:      0,
		Difficulty: difficulty,
	}

	genesisBlock.Hash = CalculateHash(genesisBlock)
	return genesisBlock
}
