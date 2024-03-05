package blockchain

// Take data from block

// create counter (nonce) starting at 0

// create a hash && plus counter 

// check hash meet requirement ? valid : repeat last step

import (
	"fmt"
	"bytes"
	"math"
	"math/big"
	"log"
	"encoding/binary"
	"crypto/sha256"
)

const Difficulty = int64(16)

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := int64(0)

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}


func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.prevHash,
			pow.block.data,
			ToHex(nonce),
			ToHex(Difficulty),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.block.nonce)
	hash := sha256.Sum256(data)

	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}