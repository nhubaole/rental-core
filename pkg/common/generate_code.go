package common

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"
)


var b = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

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
func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(b))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, b)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
