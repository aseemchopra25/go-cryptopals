package basics

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBytes(h string) []byte {
	// we convert the string to byte array
	hb := []byte(h)
	// calculate length of decoded hex first
	// to create an empty byte array storing the decoded output
	dst := make([]byte, hex.DecodedLen(len(hb)))
	// decode the hex to bytes
	hex.Decode(dst, hb)
	// now dst contains hex decoded to bytes
	return dst
}
func HexToB64(h string) string {
	dst := HexToBytes(h)

	// Convert these decoded bytes to base64
	// Concept: base64 clumps 8 bit bytes together to form 24 bit chunks
	// these 24 bit chunks are divided into 6 bit base64 digits
	// = is a padding character for base64

	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(dst)))
	base64.StdEncoding.Encode(b64, dst)
	return string(b64)
}

func FixedXor(a string, b string) string {
	a1 := HexToBytes(a)
	b1 := HexToBytes(b)
	ret := make([]byte, len(a1))
	for i := 0; i < len(a1); i++ {
		ret[i] = a1[i] ^ b1[i]
	}
	return hex.EncodeToString(ret)

}
