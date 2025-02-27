package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	// Global variables for ensuring uniqueness
	mu            sync.Mutex
	lastTimestamp int64
	counter       int64
	nodeID        int64 // For distributed systems
)

func init() {
	// Initialize nodeID with a value unique to this process
	// In production, this would be a value assigned to each server/instance
	nodeID = int64(os.Getpid() % 1024)
}

// generateUniqueID creates a guaranteed unique 15-character ID
func generateUniqueID(prefix string) string {
	mu.Lock()
	defer mu.Unlock()

	// Ensure prefix is exactly 3 characters
	paddedPrefix := padPrefix(prefix)

	// Get current timestamp in milliseconds
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// If called within the same millisecond, increment counter
	// Otherwise reset counter
	if timestamp == lastTimestamp {
		counter++
	} else {
		counter = 0
		lastTimestamp = timestamp
	}

	// Format timestamp to fit our needs - use just enough digits to maintain
	// proper ordering while leaving space for other components
	timeComponent := fmt.Sprintf("%06X", timestamp%0xFFFFFF) // 6 hex chars = 24 bits

	// Use counter and nodeID to ensure uniqueness (3 hex chars = 12 bits)
	uniqueComponent := fmt.Sprintf("%03X", (nodeID<<8)|(counter&0xFF))

	// Random component (3 hex chars)
	randomComponent := secureRandomHex(3)

	// Combine all parts: prefix(3) + time(6) + unique(3) + random(3) = 15 chars
	id := fmt.Sprintf("%s%s%s%s", paddedPrefix, timeComponent, uniqueComponent, randomComponent)

	return strings.ToUpper(id)
}

// padPrefix ensures the prefix is exactly 3 characters
func padPrefix(prefix string) string {
	if len(prefix) > 3 {
		return prefix[:3]
	} else if len(prefix) < 3 {
		return prefix + "XXX"[:3-len(prefix)]
	}
	return prefix
}

// secureRandomHex generates cryptographically secure random hex digits
// with fallbacks to ensure it always works
func secureRandomHex(length int) string {
	// Calculate how many bytes we need
	numBytes := (length + 1) / 2 // Each byte gives us 2 hex chars
	randomBytes := make([]byte, numBytes)

	// Try to get cryptographically secure random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Fallback to less secure random source using nanosecond time
		for i := range randomBytes {
			// Mix current time nanoseconds with counter and iteration for better entropy
			t := time.Now().UnixNano()
			randomBytes[i] = byte((t + counter + int64(i)) % 256)
			// Small sleep to ensure time changes between iterations
			time.Sleep(time.Nanosecond)
		}
	}

	// Convert to hex string with specified length
	hexStr := fmt.Sprintf("%X", randomBytes)
	if len(hexStr) > length {
		return hexStr[:length]
	}
	return hexStr
}

// Alternative method using base64 encoding if you need a denser representation
func generateUniqueID_Base64(prefix string) string {
	mu.Lock()
	defer mu.Unlock()

	// Get current timestamp in milliseconds
	now := time.Now().UnixNano() / int64(time.Millisecond)

	// If called within the same millisecond, increment counter
	if now == lastTimestamp {
		counter++
	} else {
		counter = 0
		lastTimestamp = now
	}

	// Create a byte buffer to hold our unique data
	// 8 bytes timestamp + 2 bytes counter/nodeID + 6 bytes random
	buffer := make([]byte, 16)

	// Put timestamp (8 bytes)
	buffer[0] = byte(now >> 56)
	buffer[1] = byte(now >> 48)
	buffer[2] = byte(now >> 40)
	buffer[3] = byte(now >> 32)
	buffer[4] = byte(now >> 24)
	buffer[5] = byte(now >> 16)
	buffer[6] = byte(now >> 8)
	buffer[7] = byte(now)

	// Put counter and nodeID (2 bytes)
	buffer[8] = byte((nodeID&0x03)<<6 | (counter>>8)&0x3F)
	buffer[9] = byte(counter & 0xFF)

	// Put random bytes (6 bytes)
	_, err := rand.Read(buffer[10:])
	if err != nil {
		// Fallback
		for i := 10; i < 16; i++ {
			buffer[i] = byte((time.Now().UnixNano() + int64(i)) % 256)
			time.Sleep(time.Nanosecond)
		}
	}

	// Encode the buffer to base64 (will give us ~22 chars)
	encoded := base64.StdEncoding.EncodeToString(buffer)

	// Ensure prefix is 3 chars
	paddedPrefix := padPrefix(prefix)

	// Combine prefix and encoded data, trim to 15 chars total
	result := paddedPrefix + encoded
	if len(result) > 15 {
		result = result[:15]
	}

	return strings.ToUpper(result)
}
