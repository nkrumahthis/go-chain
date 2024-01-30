package main

import (
	"os"

	"github.com/nkrumahthis/go-chain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()

	// w := wallet.MakeWallet()
	// w.Address()
}
