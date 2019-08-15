package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Blockchain is a series of validated Blocks
var Blockchain []Block

// IsBlockValid makes sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// CalculateHash hashes using SHA256 algorithm
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + string(block.Vote.Candidate) + string(block.Vote.Voter.NationalID) + string(block.Vote.Voter.Timestamp) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock creates a new block using previous block's hash
func GenerateBlock(oldBlock Block, vote Vote) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Vote.Candidate = vote.Candidate
	newBlock.Vote.Voter.NationalID = vote.Voter.NationalID
	newBlock.Vote.Voter.Timestamp = vote.Voter.Timestamp
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}
