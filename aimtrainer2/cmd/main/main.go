package main

import (
	"fmt"

	"github.com/aidenfoxivey/aimtrainer2/pkg/merkle"
)

func main() {
	testStr := []byte("Alisya is my beloved gf")
	rootHash, _ := merkle.HashListFromBytes(testStr, 4)
	fmt.Printf("%x\n", rootHash.TopHash)
}
