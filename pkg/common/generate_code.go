package common

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"
)



func GenerateCode(prefix string) string {
	// Seed random generator
	rand.Seed(time.Now().UnixNano())
	
	// Generate a random integer and convert it to string
	randomID := strconv.Itoa(rand.Intn(1000000)) // Generates a random number up to 6 digits
	
	// Format the code as: prefix + random ID
	return prefix + randomID
}
func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func StringToAsciiNumbers(s string) string {
	var result string
	for _, r := range s {
		if r <= utf8.RuneSelf { // Check if rune fits in a single byte (ASCII)
			result = result + strconv.Itoa(int(byte(r)))
		}
	}
	return string(result)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

