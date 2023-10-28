package consensus

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	Height     int64
	Timestamp  int64
	PrevHash   string
	Data       string
	Nonce      int64
	Difficulty int
	Validator  string
	Hash       string
}

func CalculateHash(block Block) string {
	record := strconv.FormatInt(block.Height, 10) + strconv.FormatInt(block.Timestamp, 10) + block.PrevHash +
		strconv.FormatInt(block.Nonce, 10) + block.Validator + block.Data

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}
