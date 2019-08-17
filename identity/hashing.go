package identity

import (
	"crypto/sha256"
	"encoding/hex"
)

type bytes []byte

// ConvertToByte takes a string and converts it to an array of bytes
func ConvertToByte(payload string) []byte {
	return []byte(payload)
}

// BytesPadding adds padding to the bytes
func BytesPadding(b bytes) []byte {
	const padding = 0
	payload := b
	if len(payload) < 128 {
		for len(payload) <= 128 {
			payload = append(payload, byte(padding))
		}
	}
	return payload
}

// Hash computes a SHA256 hash from a given value
func Hash(value string) string {
	bytes := ConvertToByte(value)
	paddedBytes := BytesPadding(bytes)
	h := sha256.New()
	h.Write(paddedBytes)
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}
