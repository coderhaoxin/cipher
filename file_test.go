package main

import (
	"testing"
)

func TestFile(t *testing.T) {
	data := "hello, haoxin"

	write("./test.txt", data)
	file := bytes2string(read("./test.txt"))

	t.Logf("\nfile data: %s \n", file)

	if file != data {
		t.Fatal()
	}
}
