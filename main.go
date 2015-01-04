package main

import "github.com/docopt/docopt-go"
import . "github.com/tj/go-debug"
import "path"
import "fmt"
import "os"

var debug = Debug("cipher")

func main() {
	usage := `
    Usage:
      cipher <filepath> [-r] [-d]
      cipher --help
      cipher --version

    Options:
      -f=<filepath> Required, filepath to cipher
      -d            Decrypt the file
      -r            Replace file
      --help        Show this screen
      --version     Show version
  `

	args, _ := docopt.Parse(usage, os.Args[1:], true, "v1.0.0", false)

	filepath := args["<filepath>"].(string)
	replace := args["-r"].(bool)
	toDecrypt := args["-d"].(bool)

	var key string

	debug("filepath: %s, replace: %v, decrypt: %v", filepath, replace, toDecrypt)

	fmt.Println("input the key \n")

	var filedata []byte
	var result string

	// get key
	pass := readline()
	debug("input key: %s", pass)
	if pass == "" {
		// default key
		key = path.Base(filepath)
	} else {
		key = pass
	}

	filedata = read(filepath)

	if toDecrypt {
		// decrypt
		result = decrypt(filedata, key)
		debug("to decrypt - key: %s, filedata: %s", key, bytes2string(filedata))
		fmt.Println("\n" + result)
	} else {
		// encrypt
		debug("to encrypt - key: %s, filedata: %s", key, bytes2string(filedata))
		result = encrypt(filedata, key)
		fmt.Println("\n" + result)
	}

	if replace {
		fmt.Printf("\nreplace the file: %s with above data (y or n) \n", filepath)
	} else {
		os.Exit(0)
	}

	if readline() == "y" {
		// write file
		if toDecrypt {
			debug("write file - decrypt")
			write(filepath, result)
		} else {
			debug("write file - encrypt")
			write(filepath, result)
		}
	}

	os.Exit(0)
}
