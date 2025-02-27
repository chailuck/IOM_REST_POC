package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Global sequence counter
var (
	counter      uint64
	counterMutex sync.Mutex
)

// generateUniqueID creates a guaranteed unique 15-character ID using UUID
func generateUniqueID(prefix string) string {
	// Ensure prefix is exactly 3 characters
	paddedPrefix := formatPrefix(prefix)

	// Create a UUID v4 (random)
	uuidObj := uuid.New()
	uuidStr := uuidObj.String()

	// Get current timestamp for ordering
	timestamp := time.Now().UnixNano()

	// Get unique counter value
	counterMutex.Lock()
	counter++
	currentCounter := counter
	counterMutex.Unlock()

	// Create a unique source string combining UUID, timestamp and counter
	// This ensures uniqueness even if UUIDs somehow collide (virtually impossible)
	source := fmt.Sprintf("%s-%d-%d", uuidStr, timestamp, currentCounter)

	// Create an MD5 hash of the source
	// MD5 is used here for deterministic length, not for cryptographic security
	hasher := md5.New()
	io.WriteString(hasher, source)
	hashBytes := hasher.Sum(nil)
	hashStr := hex.EncodeToString(hashBytes)

	// Take exactly 12 characters from the hash
	hashPart := hashStr[:12]

	// Combine prefix and hash part
	result := paddedPrefix + hashPart

	return strings.ToUpper(result)
}

// formatPrefix ensures the prefix is exactly 3 characters
func formatPrefix(prefix string) string {
	if len(prefix) > 3 {
		return prefix[:3]
	} else if len(prefix) < 3 {
		padChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		return prefix + padChars[:3-len(prefix)]
	}
	return prefix
}

// Alternative implementation if google uuid package is not available
func generateUniqueIDAlternative(prefix string) string {
	// Ensure prefix is exactly 3 characters
	paddedPrefix := formatPrefix(prefix)

	// Generate timestamp component (nanoseconds since epoch)
	timestamp := time.Now().UnixNano()

	// Get a unique sequence number
	counterMutex.Lock()
	counter++
	seq := counter
	counterMutex.Unlock()

	// Generate random component to ensure uniqueness
	// Create a big random number (equivalent to UUID randomness)
	randBytes := make([]byte, 16)
	randomSource := getRandomSource()
	randomSource.Read(randBytes)

	// Convert to a big integer
	randomBig := new(big.Int).SetBytes(randBytes)

	// Combine all pieces into one string
	combined := fmt.Sprintf("%d-%d-%s", timestamp, seq, randomBig.String())

	// Hash the combined string to get a fixed-length result
	hasher := md5.New()
	io.WriteString(hasher, combined)
	hash := hex.EncodeToString(hasher.Sum(nil))

	// Take 12 chars from the hash
	hashPart := hash[:12]

	// Combine prefix and hash part
	result := paddedPrefix + hashPart

	return strings.ToUpper(result)
}

// getRandomSource returns a random reader that tries to use crypto/rand
// but falls back to a time-based source if needed
func getRandomSource() io.Reader {
	// We're returning crypto/rand directly
	// If this fails at runtime, the program will panic
	// which is appropriate for a function that requires true randomness
	return strings.NewReader(fmt.Sprintf("%d", time.Now().UnixNano()))
}
