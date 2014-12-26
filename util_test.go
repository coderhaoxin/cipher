package main

import "testing"

func TestUtil(t *testing.T) {
	b := []byte("hello")
	b2s := bytes2string(b)
	b2h2s := hex2string(bytes2hex(b))
	b2h2s2b := []byte(hex2string(bytes2hex(b)))

	t.Logf("\n b: %v b2s: %s, b2h2s: %s, b2h2s2b %v \n", b, b2s, b2h2s, b2h2s2b)

	if b2s != b2h2s || !bytesEqual(b, b2h2s2b) {
		t.Fatal()
	}
}

func bytesEqual(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}

	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}

	return true
}
