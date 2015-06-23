package main

import "github.com/pkg4go/pkgs/console"
import "github.com/docopt/docopt-go"
import . "github.com/tj/go-debug"
import "path"
import "fmt"
import "os"

var debug = Debug("cipher")

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

	args, _ := docopt.Parse(usage, os.Args[1:], true, "v1.2.0", false)

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
		fmt.Println(" $ input the key \n")
		key = console.InterceptLine()
	} else {
		// default key
		key = path.Base(filepath)
	}
	debug("filepath: %s, replace: %v, decrypt: %v, setKey: %v, key: %s", filepath, replace, toDecrypt, setKey, key)

	filedata = read(filepath)

	if toDecrypt {
		// decrypt
		debug("to decrypt - filedata: %s", bytes2string(filedata))
		result = decrypt(filedata, key)
	} else {
		// encrypt
		debug("to encrypt - filedata: %s", bytes2string(filedata))
		result = encrypt(filedata, key)
	}

	if interactive {
		fmt.Println(result)
	}

	if replace {
		if interactive {
			fmt.Printf("\n $ replace the file: %s with above data (y or n) \n", filepath)
			if console.ReadChar() == "y" {
				// write file
				write(filepath, result)
			}
		} else {
			// write file
			write(filepath, result)
		}
	}

	os.Exit(0)
}
