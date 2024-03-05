package main

import (
	"go-blockchain/blockchain"
)



func main () {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	blockchain.ExploreBlockChain(chain)
}