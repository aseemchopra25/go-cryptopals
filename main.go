package main

import (
	"fmt"

	"github.com/aseemchopra25/go-cryptopals/basics"
)

func main() {
	// Challenge Set 1: Basics

	fmt.Println("1. Hex to Base64")
	h := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(basics.HexToB64(h))

	fmt.Println("2. Fixed XOR")
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	fmt.Println(basics.FixedXor(a, b))

	fmt.Println("3. Single Byte XOR Cipher")
	x := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	fmt.Println(basics.SingleByteXorCipher(x))
}
