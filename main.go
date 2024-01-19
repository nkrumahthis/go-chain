package main

import (
	"fmt"

	"github.com/nkrumahthis/go-chain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block")
	chain.AddBlock("another block")

	for _, block := range chain.Blocks {
		fmt.Println("---block")
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("prev: %x\n", block.PrevHash)
		fmt.Println("*")
	}
}
