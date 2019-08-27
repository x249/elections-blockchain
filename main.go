package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/phr3nzy/elections-blockchain/identity"
	"github.com/phr3nzy/elections-blockchain/p2p"

	"github.com/joho/godotenv"
	"github.com/phr3nzy/elections-blockchain/core"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := core.Block{}
		genesisBlock = core.Block{
			Index:     0,
			Timestamp: t.String(),
			Vote: core.Vote{
				Candidate:     "Osama",
				VoterIdentity: identity.GenerateBlockchainIdentity(identity.VoterInfo{Name: "Osama", NationalID: "0000-0000-0000-0000"}),
			},
			Hash:     core.CalculateBlockHash(genesisBlock),
			PrevHash: "",
		}
		log.Printf("Genesis block created successfully with hash: \n\x1b[32m{%v}\x1b[0m", genesisBlock.Hash)

		core.Mutex.Lock()
		core.Blockchain = append(core.Blockchain, genesisBlock)
		core.Mutex.Unlock()
	}()

	peerPort, err := strconv.Atoi(os.Getenv("PEER_PORT"))
	if err != nil {
		log.Fatalf("%v", err)
	}

	go p2p.Start(peerPort, 0)
	log.Fatal(core.Run())
}
