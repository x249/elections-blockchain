package core

// Vote represents a 'vote' casted
type Vote struct {
	Candidate     string
	VoterIdentity string
}

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Vote      Vote
	Hash      string
	PrevHash  string
}
