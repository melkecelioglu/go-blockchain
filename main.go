package main
import (
	"fmt"
	"time"
	"crypto/sha256"

)

type Block struct {
	timestamp time.Time
	transactions []string
	prevhash	 []byte
	Hash		 []byte
}

func main(){
	abc := []string{" A sent 50 coins to BC"}
	xyz := Blocks(abc, []byte{})
	fmt.Println("this is our first block")
	Print(xyz)


	pqrs := []string{ "PQ sent 230 coins to RS"}
	klmn := Blocks(pqrs, xyz.Hash)
	fmt.Println("this is our second block")
	Print(klmn)
}

func Blocks(transactions []string, prevhash []byte) *Block {
	currentTime:= time.Now()
	return &Block{
		timestamp: currentTime,
		transactions: transactions,
		prevhash: prevhash,
		Hash: NewHash(currentTime, transactions, prevhash),
	}
}

func NewHash(time time.Time, transactions[]string, prevhash[]byte) []byte{
	input := append(prevhash, time.String()...)
	for transaction:= range transactions {
		input = append(input, string(rune(transaction))...)

	}
	hash:= sha256.Sum256(input)
	return hash[:]
}


func Print(block *Block){
	fmt.Println("time: ", block.timestamp.String())
	fmt.Println("prevhash:" , block.prevhash)
	fmt.Println("Hash", block.Hash)
	Transaction(block)
}

func Transaction(block *Block){
	fmt.Println("Transactions:")
	for i, transaction :=range block.transactions{
		fmt.Printf("%v: %q", i, transaction)
		}
	}