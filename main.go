package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/nkrumahthis/go-chain/blockchain"
)

type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()
	for {
		block := iter.Next()

		fmt.Println("---block")
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("prev: %x\n", block.PrevHash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("*")
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}

	}
}

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block")
	chain.AddBlock("another block")
}
