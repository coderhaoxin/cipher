package main

import "testing"

func TestCipher(t *testing.T) {
	s := "hello"
	k := "haoxin"

	result := decode([]byte(encode([]byte(s), k)), k)

	t.Logf("\n result: %s \n", result)

	if result != s {
		t.Fatal()
	}
}
