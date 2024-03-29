package utils

// Credit: ChatGPT
// The code below is fully generated by AI

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateAESKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if length < unpadding {
		return nil, fmt.Errorf("decrypting error! Maybe wrong key")
	}

	return src[:(length - unpadding)], nil
}

func DecryptCBC(key []byte, encrypted string) ([]byte, error) {
	parts := strings.Split(encrypted, ".")

	ciphertext, err := hex.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}

	iv, err := hex.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext
	mode := cipher.NewCBCDecrypter(block, iv)
	decryptedText := make([]byte, len(ciphertext))
	mode.CryptBlocks(decryptedText, ciphertext)

	decryptedText, err = PKCS5UnPadding(decryptedText)

	if err != nil {
		return nil, err
	}

	return decryptedText, nil
}
