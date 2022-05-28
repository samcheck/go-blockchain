package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/samcheck/go-blockchain/server"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	port := flag.Uint("port", 5000, "TCP port for blockchain server")

	flag.Parse()
	fmt.Println(*port)
	appServer := server.NewBlockChainServer(uint16(*port))
	appServer.Run()
	// app := NewBlockChainServer(uint16(*port))
}
