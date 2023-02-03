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
    Hmac          EncoderType = "hmac"   // use hmac hash
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
    // Hash Generate and return a hash value in []byte format
    Hash(src string) ([]byte, error)
}
```

1. After using the `**New()**` method to obtain an Encoder instance, you can call its `**Encode()**` method to obtain the hash result of the specified data.
2. use the `**Verify()**` method to compares a encoded data with its possible plaintext equivalent.
3. `GetSalt()` method can get the automatically generated random string salt or the salt provided by yourself.
4. When you call the `**New()**` method, you can also provide zero or more `**Options**` to help initialize an `**Encoder**` instance.
5. `**Hash()**` method returns the unencoded hash value in `**[]byte**` format
6. `**Verify()**` method can only verify the result returned by `**Encode()**` method, not `**Hash()**` method. The result returned by `**Hash()**` is handled by the user.
### Installation

```bash
go get gopkg.in/encoder.v1
```

Run `go test` in the package's directory to run test case.

### Usage

Following is an example of the usage of this package:
---
#### Argon2
> `Argon2` [WIKI](https://en.wikipedia.org/wiki/Argon2)  
> Argon2 is a key derivation function that was selected as the winner of the 2015 Password Hashing Competition.  

> hash result example: $argon2id$v=19$m=65536,t=1,p=4$Te27FDFdjc7lofyIqKc4FA$XLkAG/lwiZGVvq5jVMTAIgqV2NZGDXSKFvKoIuCx/Pc
##### use default options

```go
package main

import (
	"encoding/base64"
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
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
	"encoding/base64"
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
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
> `Bcrypt` [WIKI](https://en.wikipedia.org/wiki/Bcrypt)  
> Besides incorporating a salt to protect against rainbow table attacks, bcrypt is an adaptive function: over time, the iteration count can be increased to make it slower, so it remains resistant to brute-force search attacks even with increasing computation power.  

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
> `PBKDF2` (Password-Based Key Derivation Function 2) [WIKI](https://en.wikipedia.org/wiki/PBKDF2)

> PBKDF2 applies a pseudorandom function, such as hash-based message authentication code (HMAC), to the input password or passphrase along with a salt value and repeats the process many times to produce a derived key, which can then be used as a cryptographic key in subsequent operations. The added computational work makes password cracking much more difficult, and is known as key stretching.  

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
	"encoding/base64"
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
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
	"encoding/base64"
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
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
#### Scrypt
> `Scrypt` [WIKI](https://en.wikipedia.org/wiki/Scrypt)  
> scrypt is a cryptographic derivation algorithm created by Colin Percival. Using the scrypt algorithm to generate derived keys requires a lot of memory. The scrypt algorithm was published as the RFC 7914 standard in 2016.  
> The main function of the password derivation algorithm is to generate a series of derivative passwords according to the initialized master password. This algorithm is mainly to resist the attack of brute force cracking. By increasing the complexity of password generation, it also increases the difficulty of brute force cracking  

> hash result example: d3bec465e05605b71678ba83be09c602ff589c78d3f2edb9fb0e83a9f34fd81f

##### use default options
```go
package main

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set cost
	encoding := encoder.New(types.Scrypt)
	// encoding := encoder.NewScryptEncoder()

	hash, err := encoding.Encode(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
zero or more options can be used each time， supported options for `scrypt`：
* `WithSaltLen` Length of the random salt. default value is 16
* `WithSalt`   Specify the salt, do not use the automatically generated salt, default automatically generate random strings
* `WithN` CPU/memory cost parameter, default value is 32768
* `WithR` block size parameter, default value is 8
* `WithP` parallelisation parameter, default value is 8
```go
package main

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/scrypt"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set cost
	encoding := encoder.New(types.Scrypt, scrypt.WithSaltLen(32))
	// encoder.NewScryptEncoder(scrypt.WithSaltLen(32))

	hash, err := encoding.Encode(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
#### Hkdf
> `Hkdf` HMAC-based KDF(key derivation function) [WIKI](https://en.wikipedia.org/wiki/HKDF)  
> HKDF is a simple key derivation function (KDF) based on the HMAC message authentication code. It was initially proposed by its authors as a building block in various protocols and applications, as well as to discourage the proliferation of multiple KDF mechanisms. The main approach HKDF follows is the "extract-then-expand" paradigm, where the KDF logically consists of two modules: the first stage takes the input keying material and "extracts" from it a fixed-length pseudorandom key, and then the second stage "expands" this key into several additional pseudorandom keys (the output of the KDF).   
> It can be used, for example, to convert shared secrets exchanged via Diffie–Hellman into key material suitable for use in encryption, integrity checking or authentication

> hash result example: 8477efcdc326dd85fb6713e9bd6a44fba66917a5d9e890a657bed2d7c6294c01

##### use default options
```go
package main

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set cost
	encoding := encoder.New(types.Hkdf)
	// encoder.NewHkdfEncoder()

	hash, err := encoding.Encode(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
zero or more options can be used each time， supported options for `scrypt`：
* `WithSaltLen` Length of the random salt. default value is 16
* `WithSalt`   Specify the salt, do not use the automatically generated salt, default automatically generate random strings
* `WithHashLen` The length of the generated hash, default value is 32
* `WithInfo` calling HMAC as the message field, default value is ""
* `WithHasFunc` Hash function used this time, default value is `sha256.New`
```go
package main

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/hkdf"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set cost
	encoding := encoder.New(types.Hkdf, hkdf.WithSaltLen(32), hkdf.WithHashLen(64))
	// encoder.NewHkdfEncoder(hkdf.WithSaltLen(32), hkdf.WithHashLen(64))

	hash, err := encoding.Encode(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
	salt, _ := encoding.GetSalt()
	fmt.Println("salt: ", base64.StdEncoding.EncodeToString(salt))
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
#### HMAC
> `hmac` Keyed-Hashing for Message Authentication [WIKI](https://en.wikipedia.org/wiki/HMAC)    
> In cryptography, an HMAC (sometimes expanded as either keyed-hash message authentication code or hash-based message authentication code) is a specific type of message authentication code (MAC) involving a cryptographic hash function and a secret cryptographic key. As with any MAC, it may be used to simultaneously verify both the data integrity and authenticity of a message.  
> HMAC can provide authentication using a shared secret instead of using digital signatures with asymmetric cryptography. It trades off the need for a complex public key infrastructure by delegating the key exchange to the communicating parties, who are responsible for establishing and using a trusted channel to agree on the key prior to communication.

> hash result example: 1bc6681912e7162213dd29ffa4518300b5a19496c181ebc7d694b40f679ed5eb

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
	// set cost
	encoding := encoder.New(types.Hmac)
	// encoder.NewHmacEncoder()

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
##### use custom options
zero or more options can be used each time， supported options for `hmac`：
* `WithKey` is the secret key. default value is 16
* `WithHasFunc` Hash function used this time, default value is `sha256.New`
```go
package main

import (
	"crypto/sha512"
	"fmt"
	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/hmac"
	"gopkg.in/encoder.v1/types"
)

func main() {
	data := "hello world"
	// set cost
	encoding := encoder.New(types.Hkdf, hmac.WithKey("my secrets"), hmac.WithHasFunc(sha512.New))
	// encoder.NewHmacEncoder(hmac.WithKey("my secrets"), hmac.WithHasFunc(sha512.New))

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