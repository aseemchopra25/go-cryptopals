package basics

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToB64(h string) string {
	// we convert the string to byte array
	hb := []byte(h)
	// calculate length of decoded hex first
	// to create an empty byte array storing the decoded output
	dst := make([]byte, hex.DecodedLen(len(hb)))
	// decode the hex to bytes
	hex.Decode(dst, hb)
	// now dst contains hex decoded to bytes

	// to convert these decoded bytes to base64
	// concept: base64 clumbs 8 bit bytes together to form 24 bit chunks
	// these 24 bit chunks are divided into 6 bit base64 digits
	// = is a padding character for base64
	// which you must've often seen

	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(dst)))
	base64.StdEncoding.Encode(b64, dst)
	return string(b64)
}
