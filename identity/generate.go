package identity

import "github.com/phr3nzy/elections-blockchain/hashing"

// VoterInfo stores all info regarding a voter
type VoterInfo struct {
	NationalID string
	Name       string
}

// GenerateBlockchainIdentity will create a unique identifier for the voter
func GenerateBlockchainIdentity(payload VoterInfo) string {
	IDInfo := payload
	Hash := hashing.Hash
	return Hash(Hash(IDInfo.NationalID) + Hash(IDInfo.Name))
}
