package main

import (
	"fmt"
	"strconv"

	"github.com/Mousticke/go-blockchain-simple/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send Crypto to Mousticke")
	bc.AddBlock("Send 2 more Crypto to Mousticke")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", string(block.Data))
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}
}
