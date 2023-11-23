package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func Encrypt(plaintext string) (string, error) {
	encryptionKey := os.Getenv("CIPHER_KEY")
	key := adjustKey(encryptionKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Convert the plaintext to a byte slice
	plaintextBytes := []byte(plaintext)

	// Pad the plaintext to a multiple of the block size
	padLen := aes.BlockSize - len(plaintextBytes)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	plaintextBytes = append(plaintextBytes, padding...)

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintextBytes))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintextBytes)

	// Combine the IV and the ciphertext into a single byte slice
	encryptedData := append(iv, ciphertext...)

	// Convert the encrypted data to a hexadecimal string
	encryptedText := hex.EncodeToString(encryptedData)
	return encryptedText, nil
}

func Decrypt(encryptedText string) (string, error) {
	encryptionKey := os.Getenv("CIPHER_KEY")

	// Decode the encrypted text from hex
	encryptedBytes, err := hex.DecodeString(encryptedText)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		os.Exit(1)
	}

	// Adjust the encryption key to be 32 bytes long
	adjustedKey := adjustKey(encryptionKey)

	// Split the encrypted data into the IV and ciphertext
	iv := encryptedBytes[:aes.BlockSize]
	ciphertext := encryptedBytes[aes.BlockSize:]

	// Create a new AES cipher block
	block, err := aes.NewCipher(adjustedKey)
	if err != nil {
		fmt.Println("Error creating AES cipher:", err)
		os.Exit(1)
	}

	// Create a cipher mode for AES-CBC
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the ciphertext
	decryptedData := make([]byte, len(ciphertext))
	mode.CryptBlocks(decryptedData, ciphertext)

	// Remove padding from the decrypted data if necessary
	decryptedData = removePadding(decryptedData)

	// Print the decrypted data
	return string(decryptedData), nil
}

// Function to adjust the key length to 32 bytes by padding or truncating
func adjustKey(key string) []byte {
	keyBytes := []byte(key)
	adjustedKey := make([]byte, 32)
	copy(adjustedKey, keyBytes)
	return adjustedKey
}

// Function to remove PKCS7 padding
func removePadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	if padding <= 0 || padding > aes.BlockSize {
		return data // Padding is invalid, return the original data
	}

	return data[:len(data)-padding]
}
