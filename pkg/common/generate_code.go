package common

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strconv"
	"unicode/utf8"
)

// user id + room id + date
// user id + bill id + date
var secret = "nhar94##coratnhieue@edat"
var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func GenerateCode(context string, key1 int, key2 int, key3 int) string {
	// Combine key1 and key2, key3 into a single integer
	combinedID := strconv.Itoa(key1) + strconv.Itoa(key2) + strconv.Itoa(key3)
	fmt.Println(combinedID )
	plainText := []byte(combinedID)
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return context + StringToAsciiNumbers(Encode(cipherText))
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
	block, err := aes.NewCipher([]byte(bytes))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
