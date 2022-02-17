# jax-vanity

[![Build Status](https://travis-ci.org/MarinX/btc-vanity.svg?branch=master)](https://travis-ci.org/MarinX/btc-vanity)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarinX/btc-vanity)](https://goreportcard.com/report/github.com/MarinX/btc-vanity)
[![GoDoc](https://godoc.org/github.com/MarinX/btc-vanity?status.svg)](https://godoc.org/github.com/MarinX/btc-vanity)
[![License MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

Generate Jaxnet Vanity Address

This is fork of http://github.com/MarinX/btc-vanity project

## What is jaxnet vanity address?
Vanity addresses are valid bitcoin addresses that contain human-readable messages. 

For example, `1JAXeBPzzD72PUXLzCkYAtGFYmK5vYNR33` is a valid address that contains the letters forming the word "JAX" as the first 3 Base-58 letters. Vanity addresses require generating and testing billions of candidate private keys, until one derives a jaxnet address with the desired pattern...[reference](https://github.com/bitcoinbook/bitcoinbook/blob/develop/ch04.asciidoc)


### Length of address
The frequency of a vanity pattern (1KidsCharity) and average time-to-find on a desktop PC

| Length | Pattern      | Frequency            | Average search time |
|--------|--------------|----------------------|---------------------|
| 1      | 1K           | 1 in 58 keys         | < 1 milliseconds    |
| 2      | 1Ki          | 1 in 3,364           | 50 milliseconds     |
| 3      | 1Kid         | 1 in 195,000         | < 2 seconds         |
| 4      | 1Kids        | 1 in 11 million      | 1 minute            |
| 5      | 1KidsC       | 1 in 656 million     | 1 hour              |
| 6      | 1KidsCh      | 1 in 38 billion      | 2 days              |
| 7      | 1KidsCha     | 1 in 2.2 trillion    | 3–4 months          |
| 8      | 1KidsChar    | 1 in 128 trillion    | 13–18 years         |
| 9      | 1KidsChari   | 1 in 7 quadrillion   | 800 years           |
| 10     | 1KidsCharit  | 1 in 400 quadrillion | 46,000 years        |
| 11     | 1KidsCharity | 1 in 23 quintillion  | 2.5 million years   |

## Using library
#### Install
```sh
go get github.com/tjaxer/jax-vanity
```

#### Use it as library

```go
package main

import (
	"fmt"

	"github.com/tjaxer/jax-vanity"
)

func main() {

	// create configuration
	cfg := &jaxvanity.Config{
		// buffered channel, more buffer, faster to find matching pattern
		Buffer: 5,
		// if you want to use testnet, set true
		TestNet: false,
	}

	btc := jaxvanity.New(cfg)

	// find a patters eg adddress which starts with "ab"
	address, err := btc.Find("ab", true)
	if err != nil {
		panic(err)
	}

	// print our custom public key
	fmt.Printf("PUBLIC KEY\n%s\n", address.PublicKey())

	// print our private key so it can be imported in most btc wallets
	fmt.Printf("PRIVATE KEY\n%s\n", address.PrivateKey())
}

```

#### Use it as a CLI tool
```sh
go get github.com/tjaxer/jax-vanity/cmd/jaxvanity
```
```sh
Usage: ./jaxvanity [OPTIONS] pattern
Example: ./jaxvanity

  -help
        Show usage message
  -testnet
        Use testnet
  -threads int
        How many threads you want to spawn (default 16)
  -uncompressed
        Use uncompressed public key 
```

## Tests
No magic, just run

```sh
go test -v
```

## Credits
- [Mastering Bitcoin by Andreas Antonopoulos](https://github.com/bitcoinbook/bitcoinbook)

## License
This library is under the MIT License
