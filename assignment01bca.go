package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

func NewBlock(transaction string, nonce int, previousHash string) Block {
	b := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	b.CurrentHash = b.CalculateHash()
	return b
}

func ListBlocks(blocks []Block) { // Update the function signature
	fmt.Println("==== Block List ====")
	for i, block := range blocks {
		fmt.Printf("Block #%d\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		if i > 0 {
			fmt.Printf("Previous Block's Hash: %s\n", blocks[i-1].CurrentHash)
		}
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println("===================")
	}
}

func (b Block) CalculateHash() string {
	stringToHash := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func ChangeBlock(block Block, newTransaction string) Block {
	block.Transaction = newTransaction
	block.CurrentHash = block.CalculateHash()
	return block
}

func VerifyChain(blocks []Block) bool {
	for i := 1; i < len(blocks); i++ {
		if blocks[i].PreviousHash != blocks[i-1].CurrentHash {
			return false
		}
	}
	return true
}
