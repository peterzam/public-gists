package main

import (
	"flag"
	"fmt"
	"log"

	"codeberg.org/peterzam/gobal/token"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	rpc := flag.String("r", "https://eth-mainnet.public.blastapi.io", "jsonrpc url")
	addr := flag.String("a", "", "wallet address")
	contract := flag.String("c", "", "contract address")

	flag.Parse()
	c, err := ethclient.Dial(*rpc)
	if err != nil {
		log.Fatal(err)
	}
	client := &token.Client{c}

	if *contract == "" {
		fmt.Println(client.GetEthBal(*addr))
	} else {
		fmt.Println(client.GetErc20Bal(*contract, *addr))
	}

}
