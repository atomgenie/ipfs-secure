package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// EncodeAES Use AES for cipher
func EncodeAES(data []byte, key [32]byte) ([]byte, error) {
	AEScipher, err := aes.NewCipher(key[:])

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(AEScipher)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	text := gcm.Seal(nonce, nonce, data, nil)

	return text, nil

}

// DecodeAES Decode AES data
func DecodeAES(data []byte, key [32]byte) ([]byte, error) {
	AEScipher, err := aes.NewCipher(key[:])

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(AEScipher)

	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	if len(data) < nonceSize {
		return nil, fmt.Errorf("Invalid data")
	}

	nonce, rawText := data[:nonceSize], data[nonceSize:]

	text, err := gcm.Open(nil, nonce, rawText, nil)

	return text, err
}
