package aes_encryption

import (
	"crypto/rand"

	"github.com/mateusmrangel/client-encryption-go/utils"
)

func GenerateCEK(bitLength int) []byte {
	byteLength := utils.ByteLength(bitLength)
	key := make([]byte, byteLength)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}
