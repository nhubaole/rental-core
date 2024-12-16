package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Float64PtrToInt64Ptr(f *float64) *int64 {
	if f == nil {
		return nil
	}
	i := int64(*f)
	return &i
}

func IntToFloat64Ptr(i int) *float64 {

	f := float64(i)
	return &f
}

func IfNullStr(requestVal, templateVal *string) string {
	if requestVal != nil && *requestVal != "" {
		return *requestVal
	}
	if templateVal != nil {
		return *templateVal
	}
	return ""
}

func IfNullFloat64(requestVal, templateVal *float64) float64 {
	if requestVal != nil {
		return *requestVal
	}
	if templateVal != nil {
		return *templateVal
	}
	return 0
}

func IfNullInt64(requestVal, templateVal *int64) int64 {
	if requestVal != nil {
		return *requestVal
	}
	if templateVal != nil {
		return *templateVal
	}
	return 0
}

func EncryptBase64AES(base64Input, key string) (string, error) {
	// Decode the input from base64
	data, err := base64.StdEncoding.DecodeString(base64Input)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %w", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptBase64AES(encryptedBase64, key string) (string, error) {
	// Decode the encrypted base64 input
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func ConvertIntToPointerInt32(input *int) *int32 {
	if input == nil {
		return nil
	}

	value := int32(*input)
	return &value
}
