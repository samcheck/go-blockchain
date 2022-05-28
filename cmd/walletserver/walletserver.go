package main

import (
	"flag"
	"fmt"
	"log"

	walletserver "github.com/samcheck/go-blockchain/wallet_server"
)

func init() {
	log.SetPrefix("Wallet Server: ")
}

func main() {

	port := flag.Uint("port", 8080, "TCP port for wallet server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain gateway")
	flag.Parse()
	fmt.Println(*port)
	app := walletserver.NewWalletServer(uint16(*port), *gateway)
	app.Run()

}
