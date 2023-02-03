package shared

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const letterBytes = "346789ABCDEFGHJKMNPQRTUVWXY"
const lenLetterBytes = int64(len(letterBytes))

// GenerateRandString ...
func GenerateRandString(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length should be greater than 0")
	}

	b := make([]byte, length)
	for i := range b {
		nBig, err := rand.Int(rand.Reader, big.NewInt(lenLetterBytes))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[nBig.Int64()]
	}
	return string(b), nil
}
