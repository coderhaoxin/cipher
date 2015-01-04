[![Build status][travis-img]][travis-url]
[![License][license-img]][license-url]
[![GoDoc][doc-img]][doc-url]

### cipher
a command line tool to encrypt text files

### usage

```sh
go get github.com/onebook/cipher

cipher --help

cipher example.txt -r    # -r: replace origin file with encrypted file

cipher example.txt -d -r # -d: decrypt
```

### License
MIT

[travis-img]: https://img.shields.io/travis/onebook/cipher.svg?style=flat-square
[travis-url]: https://travis-ci.org/onebook/cipher
[license-img]: http://img.shields.io/badge/license-MIT-green.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[doc-img]: http://img.shields.io/badge/GoDoc-reference-blue.svg?style=flat-square
[doc-url]: http://godoc.org/github.com/onebook/cipher
