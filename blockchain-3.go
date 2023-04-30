package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type BlockChainBlock struct {
	previous_block_hash string
	transaction_list    []string
	block_data          string
	block_hash          string
}

func NewBlockChainBlock(previous_block_hash string, transaction_list []string) *BlockChainBlock {
	block_data := strings.Join(transaction_list, "-") + " - " + previous_block_hash
	block_hash := fmt.Sprintf("%x", sha256.Sum256([]byte(block_data)))
	return &BlockChainBlock{previous_block_hash, transaction_list, block_data, block_hash}
}

func main() {

	fmt.Println("Name: ")
	var x string
	fmt.Scanln(&x)

	fmt.Println("Who do you want to send it to: ")
	var y string
	fmt.Scanln(&y)

	b1 := " sends 10 BTC to "

	block1 := NewBlockChainBlock("firstblock", []string{x, b1, y})
	fmt.Printf("Block 1 data: %s\n", block1.block_data)
	fmt.Printf("Block 1 hash: %s\n", block1.block_hash)
}
