package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"
	//merkletree "github.com/wealdtech/go-merkletree/v2"
)

// 定义区块结构block

type Block struct {
	PreBlockHash []byte
	Hash         []byte
	TimeStamp    int64
	Nonce        int
	//交易信息
	Data []byte
}

type BlockChain struct {
	blockchain []*Block
}

// type BlockSend struct {
// 	Block             Block
// 	preBlockProofHash []byte
// 	txProofHash       []byte
// }

// func (b *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
// 	headers := bytes.Join([][]byte{b.PreBlockHash, timestamp, b.Data}, []byte{})
// 	hash := sha256.Sum256(headers)
// 	b.Hash = hash[:]
// }

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		PreBlockHash: preBlockHash,
		Data:         []byte(data),
		TimeStamp:    time.Now().Unix(),
		Hash:         []byte{},
		Nonce:        0,
	}
	pow := NewPow(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

func (bc *BlockChain) AddBlock(data string) {
	preBlock := bc.blockchain[len(bc.blockchain)-1]
	block := NewBlock(data, preBlock.Hash)
	bc.blockchain = append(bc.blockchain, block)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
func main() {
	bc := NewBlockChain()
	bc.AddBlock("1")
	bc.AddBlock("2")
	for _, block := range bc.blockchain {
		fmt.Printf("Prev. hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewPow(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

// 序列化
func (b *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	encode.Encode(b)
	return buffer.Bytes()
}

// 反序列化
func Deserialize(b []byte) *Block {
	var block Block
	decode := gob.NewDecoder(bytes.NewReader(b))
	decode.Decode(&block)
	return &block
}

// 生成区块
// func generateNewblock() Block {
// 	//两个目标值强块和弱块

// }

// func main() {
// 	//开启网络通信

// 	//创建初始节点

// 	//开始挖矿（可用http控制）

// 	data1 := []byte{1, 2, 3, 4}
// 	data2 := []byte("hello,world")

// 	tree, err := merkletree.New([][]byte{
// 		data1,
// 		data2,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	root := tree.Root()
// 	proof, err := tree.GenerateProof(data1)

// 	if err != nil {
// 		panic(err)
// 	}

// 	verify, _ := merkletree.VerifyProof(data2, proof, root)

// 	fmt.Println(verify)

// 	// data := append(data1, data2...)
// 	// h := sha256.New()
// 	// h.Write(data)
// 	// ans := h.Sum(nil)
// 	// fmt.Println(ans)
// 	// s := hex.EncodeToString(ans)
// 	// fmt.Println(s)
// }
