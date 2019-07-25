package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash ...
func Hash(data []byte) []byte {
	dst := make([]byte, 0x40)
	src := sha256.Sum256(data)
	hex.Encode(dst, src[:])
	return dst
}
