package main

import (
	"fmt"
	a1 "github.com/LaibaImran1/assignment01bca/blockchain"
)

func main() {
	bc := a1.NewBlockchain() // Use the alias "a1" for the package

	bc.AddBlock("Alice to Bob", 1235)
	bc.AddBlock("Bob to Carol", 67890)

	bc.DisplayBlocks()

	bc.ChangeBlock(1, "New Transaction", 98765)

	if bc.VerifyChain() {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
