package main

import "github.com/docopt/docopt-go"
import . "github.com/tj/go-debug"
import "strings"
import "bufio"
import "fmt"
import "os"

// TODO: how to change blockSize from 16(default) to 32 ?

var debug = Debug("cipher")

func main() {
	usage := `
    Usage:
      cipher <filepath> [-r] [-d]
      cipher --help
      cipher --version

    Options:
      -f=<filepath> Required, filepath to cipher
      -d            Decode the file
      -r            Replace file
      --help        Show this screen
      --version     Show version
  `

	args, _ := docopt.Parse(usage, os.Args[1:], true, "v0.1.0", false)

	filepath := args["<filepath>"].(string)
	replace := args["-r"].(bool)
	toDecode := args["-d"].(bool)

	var key string

	debug("filepath: %s, replace: %v, decode: %v", filepath, replace, toDecode)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("input the key")

	var filedata []byte
	var result string

	for true {
		i, err := reader.ReadString('\n')
		input := strings.TrimSpace(i)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if key == "" {
			if input == "" {
				key = filepath // default key
			} else {
				key = input
			}

			filedata = read(filepath)

			if toDecode {
				// decode
				result = decode(filedata, key)
				debug("to decode - input: %s, key: %s, filedata: %s", input, key, bytes2string(filedata))
				fmt.Println(result)
			} else {
				// encode
				debug("to encode - input: %s, key: %s, filedata: %s", input, key, bytes2string(filedata))
				result = encode(filedata, key)
				fmt.Println(result)
			}

			fmt.Println("\n\n")

			if replace {
				fmt.Printf("replace the file: %s with above data (yes/Y/...) \n", filepath)
			} else {
				os.Exit(0)
			}
			continue
		}

		if input == "Y" || input == "yes" {
			// write file
			if toDecode {
				debug("write file - decode")
				write(filepath, result)
			} else {
				debug("write file - encode")
				write(filepath, result)
			}
		}

		os.Exit(0)
	}
}
