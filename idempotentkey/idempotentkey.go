// Package idempotentkey to be used for generating deterministic keys when working with idempotent processing
package idempotentkey

import (
	"hash/fnv"

	"github.com/google/uuid"
)

// GenerateIdempotentUUID will output a deterministic UUID based on a string input
func GenerateIdempotentUUID(seed string) uuid.UUID {
	h := fnv.New128a()
	h.Write([]byte(seed))
	strHash := h.Sum([]byte{})
	u, _ := uuid.FromBytes(strHash)
	return u
}
