package main

import (
	"fmt"
)

func lightningHash(data string) string {
	return data + "*superSeguro*"
}

type Block struct {
	Data     string
	Hash     string
	LastHash string
}

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
	genesis := Block{"gen-data", "gen-hash", "gen-lastHash"}
	return &Blockchain{Chain: []Block{genesis}}
}

func (bc *Blockchain) AddBlock(data string) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	lastHash := lastBlock.Hash
	hash := lightningHash(data + lastHash)
	block := Block{data, hash, lastHash}
	bc.Chain = append(bc.Chain, block)
}

func main() {
	blockchainPoggers := NewBlockchain()
	blockchainPoggers.AddBlock("one")
	blockchainPoggers.AddBlock("two")
	blockchainPoggers.AddBlock("three")

	for _, block := range blockchainPoggers.Chain {
		fmt.Printf("Data: %s, Hash: %s, LastHash: %s\n", block.Data, block.Hash, block.LastHash)
		fmt.Print("---------------------------------------\n")
	}
}
