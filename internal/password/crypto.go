package password

import (
	"crypto/rand"
	"encoding/binary"
)

// CryptoSource is a custom implementation of a math/rand Source so that we can use the
// math/rand package for secure random operations
type CryptoSource struct{}

// Int63 returns a cryptographically secure random int64
func (c CryptoSource) Int63() int64 {
	// Represent the int64 in bytes
	var b [8]byte
	// Read secure random bytes into a slice of our array
	rand.Read(b[:])
	// Create a uint64 from the random bytes
	tmpInt := binary.LittleEndian.Uint64(b[:])
	// Cut off signed bit so that we get a positive int64
	tmpInt = tmpInt & (1<<63 - 1)

	return int64(tmpInt)
}
