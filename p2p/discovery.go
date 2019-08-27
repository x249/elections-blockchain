package p2p

import (
	"encoding/json"
	"net/http"
	"os"
)

// Peers are all client entities in the blockchain
type Peers []Peer

// InitPeerDiscovery begins the process of identifying peers on the Blockchain
func InitPeerDiscovery() (Peers, error) {
	peerDiscoveryPool := os.Getenv("PEER_DISCOVERY_POOL_URL")
	var peers Peers
	r, err := http.Get(peerDiscoveryPool)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&peers); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return peers, nil
}
