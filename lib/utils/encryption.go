package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func (l *thing) Encrypt(passphrase, plaintext string) (cipherText string, err error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertextByte := gcm.Seal(
		nonce,
		nonce,
		[]byte(plaintext),
		nil,
	)

	return base64.StdEncoding.EncodeToString(ciphertextByte), nil
}

func (l *thing) Decrypt(passphrase, cipherText string) (plainText string, err error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return "", nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil
	}

	nonceSize := gcm.NonceSize()

	// process ciphertext
	ciphertextByte, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	nonce, ciphertextByteClean := ciphertextByte[:nonceSize], ciphertextByte[nonceSize:]

	plaintextByte, err := gcm.Open(
		nil,
		nonce,
		ciphertextByteClean,
		nil,
	)

	if err != nil {
		return "", err
	}

	return string(plaintextByte), nil
}
