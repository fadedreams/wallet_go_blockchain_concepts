package main

import (
	"fmt"
	"log"

	"github.com/fadedreams/wallet_go_blockchain_concepts/wallet"
)

func init() {
	log.SetPrefix("Blockchain Node: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
}
