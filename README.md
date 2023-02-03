# go-encoder
[![Build Status](https://app.travis-ci.com/go-encoder/encoder.svg?branch=main)](https://travis-ci.org/go-encoder/encoder)
[![GoDoc](https://godoc.org/github.com/go-encoder/encoder?status.svg)](https://pkg.go.dev/github.com/go-encoder/encoder)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-encoder/encoder)](https://goreportcard.com/report/github.com/go-encoder/encoder)

This Go package provides common interface implementation to encode a raw string password (example, during registration on a site), or later verify it (example, while login in to the site).

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

> 1. After using the `New()` method to obtain an Encoder instance, you can call its `Encode()` method to obtain the hash result of the specified data.
> 2. use the `Verify()` method to compares a encoded data with its possible plaintext equivalent.
> 3. `GetSalt()` method can get the automatically generated random string salt or the salt provided by yourself.
> 4. When you call the `New()` method, you can also provide zero or more `Options` to help initialize an `Encoder` instance.

### Installation

```bash
go get gopkg.in/encoder.v1
```

Run `go test` in the package's directory to run test case.

### Usage

Following is an example of the usage of this package:
---
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
	// types.Argon2 types.Pbkdf2 types.Bcrypt types.Hkdf types.Scrypt
 	encoding := encoder.New(types.Argon2) // or use encoder.NewArgon2Encoder()

	hash, err := encoding.Encode(data)
	if err != nil {
		return
	}
	fmt.Println(hash)
	verify, err := encoding.Verify(hash, data)
	if err != nil {
		return
	}
	if verify {
		fmt.Println("match")
    }
}
```
##### use custom options
zero or more options can be used each time， supported options for `argon2`：
* `WithMemory`    The amount of memory used by the Argon2 algorithm (in kibibytes)
* `WithTime`    The number of iterations (or passes) over the memory
* `WithThreads`    The number of threads (or lanes) used by the algorithm
* `WithSaltLen`    Length of the random salt. 16 bytes is recommended for password hashing
* `WithKeyLen`    Length of the generated key (or password hash). 16 bytes or more is recommended
* `WithSalt`   Specify the salt, do not use the automatically generated salt

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
	// encoder.NewArgon2Encoder(argon2.WithSaltLen(32), argon2.WithThreads(8))

	hash, err := encoding.Encode(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
	verify, err := encoding.Verify(hash, data)
	if err != nil {
		return
	}
	if verify {
		fmt.Println("match")
	}
}
```
