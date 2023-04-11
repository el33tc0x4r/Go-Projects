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

type Blockchain struct {
	chain []*BitcoinBlock
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	blockchain.generate_genesis_block()
	return blockchain
}

func (bc *Blockchain) generate_genesis_block() {
	bc.chain = append(bc.chain, NewBitcoinBlock("0", []string{"Genesis Block"}))
}

func (bc *Blockchain) create_block_from_transaction(transaction_list []string) {
	previous_block_hash := bc.last_block().block_hash
	bc.chain = append(bc.chain, NewBitcoinBlock(previous_block_hash, transaction_list))
}

func (bc *Blockchain) display_chain() {
	for i := 0; i < len(bc.chain); i++ {
		fmt.Printf("Data %d: %s\n", i+1, bc.chain[i].block_data)
		fmt.Printf("Hash %d: %s\n", i+1, bc.chain[i].block_hash)
		fmt.Println()
	}
}

func (bc *Blockchain) last_block() *BitcoinBlock {
	return bc.chain[len(bc.chain)-1]
}

func main() {
	t1 := "Noah sends 5 BTC to Mark"
	t2 := "Mark sends 3.4 BTC to Trevor"
	t3 := "Trevor sends 7.8 BTC to Neo"
	t4 := "Neo sends 1.1 BTC to Noah"
	t5 := "Noah sends 0.2 BTC to Diego"
	t6 := "Diego sends 0.1 BTC to Billy"

	myblockchain := NewBlockchain()

	myblockchain.create_block_from_transaction([]string{t1, t2})
	myblockchain.create_block_from_transaction([]string{t3, t4})
	myblockchain.create_block_from_transaction([]string{t5, t6})

	myblockchain.display_chain()
}
