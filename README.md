# go-encoder
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
> hash result example: $argon2id$v=19$m=65536,t=1,p=4$Te27FDFdjc7lofyIqKc4FA$XLkAG/lwiZGVvq5jVMTAIgqV2NZGDXSKFvKoIuCx/Pc
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
* `WithMemory`    The amount of memory used by the Argon2 algorithm (in kibibytes), default value is 64 * 1024
* `WithIterations`    The number of iterations (or passes) over the memory, default value is 1
* `WithThreads`    The number of threads (or lanes) used by the algorithm, default value is 4
* `WithSaltLen`    Length of the random salt. 16 bytes is recommended for password hashing, default value is 16
* `WithKeyLen`    Length of the generated key (or password hash). 16 bytes or more is recommended, default value is 32
* `WithSalt`   Specify the salt, do not use the automatically generated salt, default automatically generate random strings

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

---
#### Bcrypt
> `bcrypt` does not need to specify a salt, salt is automatically generated every time, and the `GetSalt` method will not return a salt. For the same value, the hash results of multiple executions are different, but each hash result can pass the `Verify` method, so only one result needs to be saved.

> hash result example: $2a$10$jIlPSogBSONL.Y2pO1F3H.1KmpnhKMp9npWs2Q2X9rGiE0ZetrAC2
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
 	encoding := encoder.New(types.Bcrypt) // or use encoder.NewBcryptEncoder()

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
zero or more options can be used each time， supported options for `bcrypt`：
* `WithCost` in order to adjust for greater computational power, default value is 10, min is 4 max is 31, The larger the value, the more time-consuming

```go
package main

import (
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/bcrypt"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set cost
	encoding := encoder.New(types.Bcrypt, bcrypt.WithCost(10))
	// encoder.NewBcryptEncoder(bcrypt.WithCost(10))

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

---
#### Pbkdf2
> `PBKDF2` (Password-Based Key Derivation Function 2) https://en.wikipedia.org/wiki/PBKDF2

> Its basic principle is to use a pseudo-random function (such as HMAC function, sha512, etc.), take the plaintext (password) and a salt value (salt) as an input parameter, and then repeat the operation, and finally generate the secret key.
> If the number of repetitions is large enough, the cost of cracking becomes very high. The addition of the salt value will also increase the difficulty of the "rainbow table" attack.
> User passwords are stored using the PBKDF2 algorithm, which is relatively safe.

> You can specify the `hash` function used when hashing, the default is **`sha256.New`**.
> Other hash functions available include:
> * md5.New
> * sha1.New
> * sha256.New224
> * sha256.New
> * sha512.New
> * sha512.New384
> * sha512.New512_224
> * sha512.New512_256

> hash result example: c46f9f21b89a22a36df3c1de5a91a85f5b4656f616e79d2e589b3b32a4648e80
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
	encoding := encoder.New(types.Pbkdf2) // or use encoder.NewPbkdf2Encoder()
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
zero or more options can be used each time, supported options for `pbkdf2`：
* `WithHasFunc` Hash function used this time, default value is sha256.New
* `WithKeyLen` Length of the generated key, default value is 32
* `WithIterations` The number of iterations, the more times, the longer it takes to encrypt and decrypt, default value is 10000
* `WithSaltLen` Length of the random salt, default value is 16
* `WithSalt` Specify the salt, do not use the automatically generated salt, default automatically generate random strings

```go
package main

import (
	"crypto/sha512"
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/pbkdf2"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set cost
	encoding := encoder.New(types.Pbkdf2, pbkdf2.WithHasFunc(sha512.New), pbkdf2.WithIterations(20000))
	// encoding := encoder.NewPbkdf2Encoder(pbkdf2.WithHasFunc(sha512.New), pbkdf2.WithIterations(20000))

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
