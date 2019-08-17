package identity

// VoterInfo stores all info regarding a voter
type VoterInfo struct {
	NationalID  string
	Name        string
	Fingerprint string
}

// GenerateBlockchainIdentity will create a unique identifier for the voter
func (payload VoterInfo) GenerateBlockchainIdentity() string {
	IDInfo := &payload
	return Hash(Hash((*IDInfo).NationalID) + Hash((*IDInfo).Name) + Hash((*IDInfo).Fingerprint))
}
