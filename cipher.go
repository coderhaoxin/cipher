package main

import (
	"crypto/aes"
	"crypto/cipher"
)

var common = []byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
}

const (
	keySize   = 32
	nonceLen  = 12
	blockSize = aes.BlockSize
)

func encrypt(text []byte, keys string) string {
	key, nonce := getKeyAndNonce(keys)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	result := aesgcm.Seal(nil, nonce, text, nil)
	return bytes2hex(result)
}

func decrypt(text []byte, keys string) string {
	key, nonce := getKeyAndNonce(keys)
	text = []byte(hex2string(bytes2string(text)))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	result, err := aesgcm.Open(nil, nonce, text, nil)
	if err != nil {
		panic(err)
	}

	return bytes2string(result)
}

func getKeyAndNonce(keys string) (key, nonce []byte) {
	key = []byte(keys)

	if size := len(key); size != keySize {
		if size < keySize {
			key = append(key, common[size:]...)
		} else {
			key = key[:keySize]
		}
	}

	nonce = make([]byte, nonceLen)
	copy(nonce, key[0:nonceLen])

	return
}
