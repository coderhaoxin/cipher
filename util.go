package main

import "encoding/hex"
import "bytes"

func bytes2hex(b []byte) string {
	return hex.EncodeToString(b)
}

func hex2string(h string) string {
	b, e := hex.DecodeString(h)
	if e != nil {
		return ""
	}

	return bytes2string(b)
}

func bytes2string(b []byte) string {
	return bytes.NewBuffer(b).String()
}
