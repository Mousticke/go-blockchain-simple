package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) ComputeHash() {
	mergeData := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(mergeData)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.ComputeHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.blocks = []*Block{NewGenesisBlock()}
	return bc
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send Crypto to Mousticke")
	bc.AddBlock("Send 2 more Crypto to Mousticke")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", string(block.Data))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
