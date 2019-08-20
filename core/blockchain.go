package core

import (
	"strconv"
	"time"

	"github.com/phr3nzy/elections-blockchain/hashing"
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

	if CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// CalculateBlockHash hashes using SHA256 algorithm
func CalculateBlockHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + string(block.Vote.Candidate) + string(block.Vote.VoterIdentity) + block.PrevHash
	hash := hashing.Hash
	hashed := hash(record)
	return hashed
}

// GenerateBlock creates a new block using previous block's hash
func GenerateBlock(oldBlock Block, vote Vote) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Vote.Candidate = vote.Candidate
	newBlock.Vote.VoterIdentity = vote.VoterIdentity
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateBlockHash(newBlock)

	return newBlock
}
