package utils

import (
	"math/rand"
	"time"
)

// Making the character set of random string
const charset string = "abcdefghijklmnopqrstuvwxyz0123456789"

// Function to generate a random string
func GenerateRandomString() string {
	// Generate a random time dependant string
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Creating a byte slice to store random characters
	b := make([]byte, 5)

	// Looping through the byte slice
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	// Returning the converted string from the byte slice
	return string(b)
}
