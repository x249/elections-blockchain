package p2p

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"

	net "github.com/libp2p/go-libp2p-net"
	"github.com/phr3nzy/elections-blockchain/core"
	"github.com/phr3nzy/elections-blockchain/errors"
)

// HandleStream handles incoming network streams
func HandleStream(s net.Stream) {

	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go ReadData(rw)
}

// ReadData reads incoming blockchain stream data from a channel streams
func ReadData(rw *bufio.ReadWriter) {

	for {
		str, err := rw.ReadString('\n')
		errors.HandleError(err)

		if str == "" {
			return
		}
		if str != "\n" {

			chain := make([]core.Block, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			core.Mutex.Lock()
			if len(chain) > len(core.Blockchain) {
				core.Blockchain = chain
				bytes, err := json.MarshalIndent(core.Blockchain, "", "  ")
				errors.HandleError(err)
				// Green console color: 	\x1b[32m
				// Reset console color: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			core.Mutex.Unlock()
		}
	}
}
