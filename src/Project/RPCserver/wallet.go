package main

import (
	"bytes"
	"crypto/ecdsa"
	//"crypto/elliptic"
	//"crypto/rand"
	//"crypto/sha256"
	base58 "github.com/btcsuite/btcd/btcutil/base58"
	//ripemd160 "golang.org/x/crypto/ripemd160"
	//"log"
)

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// NewWallet creates and returns a Wallet
func NewWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	pubKeyHash := base58.Decode(address)
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Equal(actualChecksum, targetChecksum)
}
