# go-encoder
[![Build Status](https://api.travis-ci.org/go-encoder/encoder.svg?branch=master)](https://travis-ci.org/go-encoder/encoder)
[![GoDoc](https://godoc.org/github.com/go-encoder/encoder?status.svg)](https://pkg.go.dev/github.com/go-encoder/encoder)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-encoder/encoder)](https://goreportcard.com/report/github.com/go-encoder/encoder)

This package in Go provides common interface implementation to encode a raw string password (example, during registration on a site), or later verify it (example, while logging in to the site).

Already supported hash types:
```go
const (
    Bcrypt        EncoderType = "bcrypt" // use bcrypt hash
    Scrypt        EncoderType = "scrypt" // use scrypt hash
    Pbkdf2        EncoderType = "pbkdf2" // use pbkdf2 hash
    Argon2        EncoderType = "argon2" // use argon2 hash
    Hkdf          EncoderType = "hkdf"   // use hkdf hash
)
```

Each hash type implements the `Encoder` interface
```go
type Encoder interface {
	// Encode returns the hash value of the given data
	Encode(src string) (string, error)
	// Verify compares a encoded data with its possible plaintext equivalent
	Verify(hash string, rawData string) (bool, error)
	// GetSalt Returns the salt if exists, otherwise nil
	GetSalt() ([]byte, error)
}
```

After using the `New()` method to obtain an Encoder instance, you can call its `Encode()` method to obtain the hash result of the specified data, 
or use the `Verify()` method to compares a encoded data with its possible plaintext equivalent.
`GetSalt()` method can get the automatically generated random string salt or the salt provided by yourself.
When you call the `New()` method, you can also provide zero or more `Options` to help initialize an `Encoder` instance.

### Installation

```bash
go get gopkg.in/encoder.v1
```

Run `go test` in the package's directory to run tests.

### Usage

Following is an example of the usage of this package:

#### Argon2
##### use default options
```go
package main

import (
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// Using the default options
	encoding := encoder.New(types.Argon2)

	hash, err := encoding.Encode(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
}
```
##### use custom options
zero or more options can be used each time， supported options for `argon2`：
* WithMemory
* WithTime
* WithThreads
* WithSaltLen
* WithKeyLen
* WithSalt

```go
package main

import (
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/argon2"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set salt length and threads
	encoding := encoder.New(types.Argon2, argon2.WithSaltLen(32), argon2.WithThreads(8))

	hash, err := encoding.Encode(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
}
```
