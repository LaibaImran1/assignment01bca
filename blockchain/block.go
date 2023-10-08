// block.go
package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents a block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// NewBlock creates a new block.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	if transaction == "" {
		panic("Transaction cannot be empty.")
	}
	if nonce < 0 {
		panic("Nonce must be a non-negative integer.")
	}
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CreateHash()
	return block
}

// CreateHash generates the hash for the block.
func (b *Block) CreateHash() {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	b.Hash = CalculateHash(data)
}

// DisplayBlock prints the block's information.
func (b *Block) DisplayBlock() {
	fmt.Printf("Transaction: %s\n", b.Transaction)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Previous Hash: %s\n", b.PreviousHash)
	fmt.Printf("Hash: %s\n", b.Hash)
	fmt.Println()
}

// CalculateHash calculates the SHA-256 hash of a given string.
func CalculateHash(stringToHash string) string {
	data := []byte(stringToHash)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
