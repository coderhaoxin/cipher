package main

import "crypto/cipher"
import "crypto/aes"

var common = []byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
	0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
}

var blockSize = aes.BlockSize
var keySize = 32

func encrypt(text []byte, keys string) string {
	key, iv := getKeyAndIV(keys)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	result := make([]byte, len(text))
	stream.XORKeyStream(result, text)

	return bytes2hex(result)
}

func decrypt(text []byte, keys string) string {
	key, iv := getKeyAndIV(keys)
	text = []byte(hex2string(bytes2string(text)))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	result := make([]byte, len(text))
	stream.XORKeyStream(result, text)

	return bytes2string(result)
}

func getKeyAndIV(keys string) (key, iv []byte) {
	key = []byte(keys)

	if size := len(key); size != keySize {
		if size < keySize {
			key = append(key, common[size:]...)
		} else {
			key = key[:keySize]
		}
	}

	iv = make([]byte, blockSize)
	copy(iv, key[0:blockSize])

	return
}
