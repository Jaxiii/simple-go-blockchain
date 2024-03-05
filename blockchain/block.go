package blockchain

import (
	"fmt"
	"bytes"
	"crypto/sha256"
	"strconv"
)

type BlockChain struct {
	blocks []*Block
}

// timestamp ?
type Block struct {
	hash		[]byte
	data		[]byte
	prevHash	[]byte
	nonce		int64
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.data, b.prevHash}, []byte{})
	hash := sha256.Sum256(info) // need a better hash :D
	b.hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.hash = hash[:]
	block.nonce = nonce

	return block
}

// smart conctract?
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis                ", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func ExploreBlockChain(chain *BlockChain) {
	for _, block := range chain.blocks {

		pow := NewProof(block)
		fmt.Printf("============================================|&&|==========================================\n#                                                                                       #\n")
		fmt.Printf("#  Data in Block: %s        					#\n", block.data)
		fmt.Printf("#  Hash: %x               #\n#                                                                                       #\n", block.hash)
		fmt.Printf("#  Previus Hash: %x       #\n", block.prevHash)
		fmt.Printf("#  Proof of Work: %s                                                                  #\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("============================================|&&|==========================================\n")
	}
}