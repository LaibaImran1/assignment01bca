//Laiba Imran
//20i-0991
//Assignment 1

// main.go
package main

import (
	a1 "github.com/LaibaImran1/assignment01bca/blockchain"
)

func main() {
	bc := a1.NewBlockchain() // Use the alias "a1" for the package

	// Change the nonce values and fix the genesis block setup
	genesisBlock := a1.NewBlock("Genesis Transaction", 0, "")
	bc.Blocks[0] = genesisBlock

	bc.AddBlock("Laiba to khadija", 12345)
	bc.AddBlock("Khadija to Ahmed", 67890)

	bc.DisplayBlocks()

	// Verify the integrity of the blockchain
	bc.VerifyChain()
}
