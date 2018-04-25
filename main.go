package main

import (
	"fmt"
	"os"
	"path"

	"github.com/docopt/docopt-go"
	"github.com/pkg4go/pkgs/console"
)

func main() {
	usage := `
  Usage:
    cipher <filepath> [-r] [-d] [-k] [-i]
    cipher --help
    cipher --version

  Options:
    -f=<filepath> Required, filepath to cipher
    -k            Set decrypt/encrypt key
    -d            Decrypt the file
    -r            Replace file
    -i            Interactive
    --help        Show this screen
    --version     Show version
`

	args, _ := docopt.Parse(usage, os.Args[1:], true, "v2.0.0", false)

	filepath := args["<filepath>"].(string)
	interactive := args["-i"].(bool)
	toDecrypt := args["-d"].(bool)
	replace := args["-r"].(bool)
	setKey := args["-k"].(bool)

	var filedata []byte
	var result string
	var key string

	// get key
	if setKey {
		fmt.Println(" $ input the key")
		key = console.InterceptLine()
	} else {
		// default key
		key = path.Base(filepath)
	}

	filedata = read(filepath)

	if toDecrypt {
		// decrypt
		result = decrypt(filedata, key)
	} else {
		// encrypt
		result = encrypt(filedata, key)
	}

	if interactive {
		fmt.Println(result)
	}

	if replace {
		if interactive {
			fmt.Printf("\n $ replace the file: %s with above data (y or n) \n", filepath)
			if console.ReadChar() == "y" {
				write(filepath, result)
			}
		} else {
			write(filepath, result)
		}
	}

	os.Exit(0)
}
