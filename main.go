package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
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
				VoterIdentity: "0000-0000-0000-0000",
			},
			Hash:     core.CalculateBlockHash(genesisBlock),
			PrevHash: "",
		}
		spew.Dump(genesisBlock)

		core.Mutex.Lock()
		core.Blockchain = append(core.Blockchain, genesisBlock)
		core.Mutex.Unlock()
	}()
	log.Fatal(core.Run())
}
