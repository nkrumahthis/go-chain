package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/nkrumahthis/go-chain/blockchain"
	"github.com/nkrumahthis/go-chain/wallet"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" balance -address ADDRESS - get the balance for an account")
	fmt.Println(" create -address ADDRESS creates a blockchain and sends genesis reward to address")
	fmt.Println(" destroy deletes the blockchain")
	fmt.Println(" print - Prints the blocks in the chain")
	fmt.Println(" send -from FROM -to TO -amount AMOUNT - Send amount of coins")
	fmt.Println(" createWallet - Creates a new wallet")
	fmt.Println(" listAddresses - Lists the adresses in our wallet file")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) createBlockChain(address string) {
	chain := blockchain.InitBlockChain(address)
	defer chain.Close()
	fmt.Println()
}

func (cli *CommandLine) getBalance(address string) {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.Close()

	balance := 0
	UTXOs := chain.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}

func (cli *CommandLine) send(from, to string, amount int) {
	chain := blockchain.ContinueBlockChain(from)
	defer chain.Close()

	tx := blockchain.NewTransaction(from, to, amount, chain)
	chain.AddBlock([]*blockchain.Transaction{tx})
	fmt.Println("Success!")
}

func (cli *CommandLine) printChain() {
	chain := blockchain.ContinueBlockChain("")
	defer chain.Close()
	iter := chain.Iterator()

	for {
		block := iter.Next()

		fmt.Println("---block")
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

func (cli *CommandLine) listAddresses() {
	wallets, err := wallet.CreateWallets()
	if err != nil {
		log.Panic(err)
	}

	addresses := wallets.GetAllAddress()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

func (cli *CommandLine) destroyChain() {
	fmt.Println("destroying chain")
	blockchain.DestroyBlockChain()
	fmt.Println("chain is destroyed")
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	balanceCmd := flag.NewFlagSet("balance", flag.ExitOnError)
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	destroyCmd := flag.NewFlagSet("destroy", flag.ExitOnError)

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listAddresses", flag.ExitOnError)

	balanceAddress := balanceCmd.String("address", "", "The address to get the balance for")
	createAddress := createCmd.String("address", "", "The address to send the genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "balance":
		err := balanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "create":
		err := createCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "destroy":
		err := destroyCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listAddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if createCmd.Parsed() {
		if *createAddress == "" {
			balanceCmd.Usage()
			runtime.Goexit()
		}
		cli.createBlockChain(*createAddress)
	}

	if balanceCmd.Parsed() {
		if *balanceAddress == "" {
			balanceCmd.Usage()
			runtime.Goexit()
		}
		cli.getBalance(*balanceAddress)
	}

	if destroyCmd.Parsed() {
		cli.destroyChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			runtime.Goexit()
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses()
	}
}

