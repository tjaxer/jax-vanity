package main

import (
	"flag"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	jaxvanity "github.com/tjaxer/jax-vanity"
	"os"
)

var buffer = flag.Int("threads", 16, "How many threads you want to spawn")
var testnet = flag.Bool("testnet", false, "Use testnet")
var uncompressed = flag.Bool("uncompressed", false, "Use uncompressed public key ")
var help = flag.Bool("help", false, "Show usage message")
var jaxtestnet = flag.Bool("jaxtestnet", false, "Create Jax network testnet address")

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

	if *testnet {
		prefix = "jax"
	}

	if *help {
		usage()
	}

	if flag.NArg() != 0 {
		prefix = flag.Arg(0)
	}


	cfg := &jaxvanity.Config{
		Buffer:  *buffer,
		TestNet: *testnet,
	}

	btc := jaxvanity.New(cfg)

	fmt.Fprintf(os.Stdout,
		"Testnet: %t\nUncompressed public: %v\nThreads: %d\nPattern: %s\nWorking...please wait\n",
		*testnet, *uncompressed, *buffer, prefix,
	)

	address, err := btc.Find(prefix, !*uncompressed )
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if !*jaxtestnet {
		fmt.Fprintf(os.Stdout, "Address: %s\n", address.Address())
		fmt.Fprintf(os.Stdout, "Public key: %s\n", address.PublicKey())
		fmt.Fprintf(os.Stdout, "Private key (WIF): %s\n", address.PrivateKey())
	} else {
		mainnetWif,_ := btcutil.DecodeWIF(address.PrivateKey())
		testnetWif,_ := btcutil.NewWIF(mainnetWif.PrivKey, &chaincfg.TestNet3Params, mainnetWif.CompressPubKey)
		var testnetPubKey  *btcutil.AddressPubKey
		if !*uncompressed {
			testnetPubKey,_ = btcutil.NewAddressPubKey(
				testnetWif.PrivKey.PubKey().SerializeCompressed(), &chaincfg.TestNet3Params)
		} else {
			testnetPubKey,_ = btcutil.NewAddressPubKey(
				testnetWif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.TestNet3Params)
		}
		fmt.Fprintf(os.Stdout, "Mainnet address: %s\n", address.Address())
		fmt.Fprintf(os.Stdout, "Testnet private key (WIF): %s\n", address.PrivateKey())
		fmt.Fprintf(os.Stdout, "Public key: %s\n", testnetPubKey.String())
		fmt.Fprintf(os.Stdout, "Testnet address: %s\n", testnetPubKey.AddressPubKeyHash().EncodeAddress())
	}
}
