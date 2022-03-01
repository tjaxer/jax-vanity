package jaxvanity

import (
	"github.com/btcsuite/btcutil"
)

// Wallet is our private/public key holder
type Wallet struct {
	privKey *btcutil.WIF
	pubKey  *btcutil.AddressPubKey
}

// IWallet is interface to our internal wallet
type IWallet interface {
	PublicKey() string
	PrivateKey() string
	Address() string
}


// Address returns encoded address
func (w Wallet) Address() string {
	return w.pubKey.AddressPubKeyHash().EncodeAddress()
}


// PublicKey returns encoded public key
func (w Wallet) PublicKey() string {
	return w.pubKey.String()
}

// PrivateKey returns private key, ready for import in most bitcoin wallets
func (w Wallet) PrivateKey() string {
	return w.privKey.String()
}
