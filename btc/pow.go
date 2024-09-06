package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

var targetBits = 24

type Pow struct {
	block  *Block
	target *big.Int
}

func NewPow(block *Block) *Pow {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &Pow{block, target}
	return pow
}

func (pow *Pow) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreBlockHash,
			IntToHex(pow.block.TimeStamp),
			pow.block.Data,
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

// IntToHex 将 int64 类型的整数转换为字节数组
func IntToHex(n int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, n)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buff.Bytes()
}

func (pow *Pow) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("%x\n", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *Pow) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
