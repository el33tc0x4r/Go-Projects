package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type BitcoinBlock struct {
	previous_block_hash string
	transaction_list    []string
	block_data          string
	block_hash          string
}

func NewBitcoinBlock(previous_block_hash string, transaction_list []string) *BitcoinBlock {
	block_data := strings.Join(transaction_list, "-") + " - " + previous_block_hash
	block_hash := fmt.Sprintf("%x", sha256.Sum256([]byte(block_data)))
	return &BitcoinBlock{previous_block_hash, transaction_list, block_data, block_hash}
}

func main() {
	t1 := "Noah sends 5 BTC to Mark"
	t2 := "Mark sends 3.4 BTC to Trevor"
	t3 := "Trevor sends 7.8 BTC to Jimmy"
	t4 := "Jimmy sends 1.1 BTC to Noah"

	block1 := NewBitcoinBlock("firstblock", []string{t1, t2})
	fmt.Printf("Block 1 data: %s\n", block1.block_data)
	fmt.Printf("Block 1 hash: %s\n", block1.block_hash)

	block2 := NewBitcoinBlock("firstblock", []string{t3, t4})
	fmt.Printf("Block 2 data: %s\n", block2.block_data)
	fmt.Printf("Block 2 hash: %s\n", block2.block_hash)
}
