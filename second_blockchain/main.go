package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)


type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte

}

type Blockchain struct{
	blocks []*Block
}


func (b *Block) DeriveHash(){
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block{
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string){
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}

func Genesis() *Block{
	return CreateBlock("Genesis", []byte{})

}

func InitBlockChain() * Blockchain{
	return &Blockchain{[]*Block{Genesis()}}
}


func main(){
	chain := InitBlockChain()
	chain.AddBlock("First block after Genesis Block")
	chain.AddBlock("Second block after Genesis Block")
	chain.AddBlock("Third block after Genesis Block")

		for _, block := range chain.blocks {
			fmt.Printf("Previous Hash: %x\n", block.PrevHash)
			fmt.Printf("Data in Block: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)	
		}
}