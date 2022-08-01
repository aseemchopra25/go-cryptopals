package main

import (
	"fmt"

	"github.com/aseemchopra25/go-cryptopals/basics"
)

func main() {
	// Challenge Set 1: Basics

	// 1. Hex to Base 64
	h := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(basics.HexToB64(h))

}
