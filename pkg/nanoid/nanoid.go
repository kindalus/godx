package nanoid

import (
	"crypto/rand"
	"math/big"
)

const (
	defaultLength   = 8
	defaultAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// ID represents a unique identifier.
type ID string

// New generates a new ID with the default length of 8 characters and the default alphabet.
// Returns an ID type instance.
func New() ID {
	return newWithLengthAndAlphabet(defaultLength, defaultAlphabet)
}

// NewWithLength generates a new ID with a custom length and the default alphabet.
// Returns an ID type instance.
func NewWithLength(length int) ID {
	return newWithLengthAndAlphabet(length, defaultAlphabet)
}

// newWithLengthAndAlphabet generates a new ID with the specified length and alphabet.
// Returns an ID type instance.
func newWithLengthAndAlphabet(length int, alphabet string) ID {
	if length <= 0 {
		length = defaultLength // fallback to default length if invalid length provided
	}
	if len(alphabet) == 0 {
		alphabet = defaultAlphabet // fallback to default alphabet if empty alphabet provided
	}

	id := make([]byte, length)
	big := big.NewInt(int64(len(alphabet)))

	for i := range id {
		idx, _ := rand.Int(rand.Reader, big)
		id[i] = alphabet[idx.Int64()]
	}

	return ID(id)
}

// String returns the ID as a string.
func (id ID) String() string {
	return string(id)
}
