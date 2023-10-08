package assignment01bca

import (
	"fmt"
)

// Blockchain represents the blockchain.
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain creates a new blockchain with the genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Transaction", 0, "")
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transaction, nonce, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// DisplayBlocks prints all the blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		block.DisplayBlock()
	}
}

// ChangeBlock changes the transaction of the given block.
func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string, newNonce int) {
	if blockIndex < 0 || blockIndex >= len(bc.Blocks) {
		panic("Invalid block index.")
	}
	block := bc.Blocks[blockIndex]
	block.Transaction = newTransaction
	block.Nonce = newNonce
	block.CreateHash()
}


// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// Recalculate the hash for the current block.
		calculatedHash := CalculateHash(fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash))

		// Verify current block's hash.
		if currentBlock.Hash != calculatedHash {
			return false
		}

		// Verify previous hash.
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
