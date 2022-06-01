package main

import (
	"flag"
	"fmt"
	"log"

	blockchainserver "github.com/samcheck/go-blockchain/internal/blockchain_server"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	port := flag.Uint("port", 5000, "TCP port for blockchain server")

	flag.Parse()
	fmt.Println(*port)
	appServer := blockchainserver.NewBlockChainServer(uint16(*port))
	appServer.Run()
	// app := NewBlockChainServer(uint16(*port))
}
