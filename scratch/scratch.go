package scratch

import (
	"fmt"

	"github.com/samcheck/go-blockchain/block"
	"github.com/samcheck/go-blockchain/wallet"
)

func main() {
	wM := wallet.NewWallet()
	wA := wallet.NewWallet()
	wB := wallet.NewWallet()
	// Wallet
	t := wallet.NewTransaction(wA.PrivateKey(), wA.PublicKey(), wA.BlockchainAddress(), wB.BlockchainAddress(), 1)
	fmt.Printf("signature %s\n", t.GenerateSignature())
	// Blockchain
	blockChain := block.NewBlockChain(wM.BlockchainAddress(), 5000)
	isAdded := blockChain.AddTransaction(wA.BlockchainAddress(), wB.BlockchainAddress(), 1, wA.PublicKey(), t.GenerateSignature())
	fmt.Println("added? ", isAdded)
	if isAdded {
		blockChain.Mining()
	}
	// blockChain.AddTransaction("Sam", "Eric", 1.0)
	// blockChain.Mining()
	// // previousHash := blockChain.LastBlock().Hash()
	// // nonce := blockChain.ProofOfWork()

	// // blockChain.CreateBlock(nonce, previousHash)
	// blockChain.Print()

	// blockChain.AddTransaction("Eric", "Sam", .50)
	// blockChain.AddTransaction("Sam", "Eric", 3.4)
	// blockChain.Mining()
	// // previousHash = blockChain.LastBlock().Hash()
	// // nonce = blockChain.ProofOfWork()
	// // blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()
	fmt.Printf("Sam %.1f\n", blockChain.CalculateTotalAmount(wA.BlockchainAddress()))
	fmt.Printf("Eric %.1f\n", blockChain.CalculateTotalAmount(wB.BlockchainAddress()))
	fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount(wM.BlockchainAddress()))
}
