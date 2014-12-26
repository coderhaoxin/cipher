package main

import "crypto/cipher"
import "crypto/aes"

var common = []byte{
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
}

var blockSize = aes.BlockSize

func encode(text []byte, keys string) string {
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

func decode(text []byte, keys string) string {
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
	// len(key) == len(iv) == blockSize (16, 24, 32)
	if size := len(key); size != blockSize {
		if size < blockSize {
			key = append(key, common[size:]...)
		} else {
			key = key[:blockSize]
		}
	}

	iv = make([]byte, blockSize)

	return
}
