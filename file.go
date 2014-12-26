package main

import "path/filepath"
import "io/ioutil"
import "os"

func read(filename string) []byte {
	path := resolve(filename)

	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return file
}

func write(filename string, data string) {
	path := resolve(filename)

	err := ioutil.WriteFile(path, []byte(data), 0644)

	if err != nil {
		panic(err)
	}
}

func resolve(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}

	cwd, err := os.Getwd()

	if err != nil {
		return filename
	}

	return filepath.Join(cwd, filename)
}
