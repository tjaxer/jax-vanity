package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MarinX/btc-vanity"
)

var buffer = flag.Int("threads", 16, "How many threads you want to spawn")
var testnet = flag.Bool("testnet", false, "Use testnet")
var help = flag.Bool("help", false, "Show usage message")

var usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] pattern\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Example: %s Kid\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	var prefix string = "JAX"

	if *help {
		usage()
	}

	if flag.NArg() != 0 {
		prefix = flag.Arg(0)
	}


	cfg := &btcvanity.Config{
		Buffer:  *buffer,
		TestNet: *testnet,
	}

	btc := btcvanity.New(cfg)

	fmt.Fprintf(os.Stdout,
		"Testnet: %t\nThreads: %d\nPattern: %s\nWorking...please wait\n",
		*testnet, *buffer, prefix,
	)

	address, err := btc.Find(prefix)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Fprintf(os.Stdout, "Address: %s\n", address.PublicKey())
	fmt.Fprintf(os.Stdout, "Public key: %s\n", address.PrivateKey())
	fmt.Fprintf(os.Stdout, "Private key: %s\n", address.PrivateKey())
	fmt.Fprintf(os.Stdout, "Wif key: %s\n", address.PrivateKey())




}
