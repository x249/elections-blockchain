package core

// Voter represents iformation regard who casted the vote and when
type Voter struct {
	NationalID string
	Timestamp  string
}

// Vote represents a 'vote' casted
type Vote struct {
	Candidate string
	Voter     Voter
}

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Vote      Vote
	Hash      string
	PrevHash  string
}
