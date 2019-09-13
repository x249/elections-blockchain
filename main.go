package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/x249/elections-blockchain/identity"
	"github.com/x249/elections-blockchain/p2p"

	"github.com/joho/godotenv"
	"github.com/x249/elections-blockchain/core"
	"github.com/x249/elections-blockchain/errors"
)

func main() {
	err := godotenv.Load()
	errors.HandleError(&err)

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
		log.Printf("Genesis block created successfully with hash: \x1b[33m%v\x1b[0m", genesisBlock.Hash)

		core.Mutex.Lock()
		core.Blockchain = append(core.Blockchain, genesisBlock)
		core.Mutex.Unlock()
	}()

	peerPort, err := strconv.Atoi(os.Getenv("PEER_PORT"))
	errors.HandleError(&err)

	go p2p.Start(peerPort, 0)
	go log.Fatal(core.Run())
}
