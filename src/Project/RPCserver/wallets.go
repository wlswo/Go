package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/crypto/ripemd160"
)
​
type Wallet struct {
	Timestamp int64             //생성시간
	Address   string            //지갑
	PrivKey   *ecdsa.PrivateKey //개인키
	PubKey    []byte            //공개키
}
​
type Wallets struct {
	Wallets map[string]*Wallet
}
​
//주소 확인
func (ws *Wallets) CheckAddr(addr string) *Wallet {
	for k, v := range ws.Wallets {
		if k == addr {
			return v
		}
	}
	return nil
}
​
//지갑 확인
func (ws *Wallets) CheckWallet(w *Wallet) bool {
	for _, v := range ws.Wallets {
		if cmp.Equal(v, w) {
			return true
		}
	}
	return false
}
​
//지갑 목록 생성
/*
func NewWallets(privKey *ecdsa.PrivateKey, pubKey []byte, addr string) *Wallets {
	return &Wallets{map[string]*Wallet{addr: NewWallet(privKey, pubKey, addr)}}
}
​*/
func NewWallets() (*Wallets, error){
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	return &wallets, err
}

//지갑 추가
func (ws *Wallets) AddWallet(privKey *ecdsa.PrivateKey, pubKey []byte,addr string) {
	w := NewWallet(privKey, pubKey, addr)
	ws.Wallets[w.Address] = w
}
​
//지갑 생성
func NewWallet(privKey *ecdsa.PrivateKey, pubKey []byte, addr string) *Wallet {
​
	//현재 시각
	time := time.Now().UTC().Unix()
​
	return &Wallet{time, addr, privKey, pubKey}
}
​
//공개 키, 개인 키 생성
func GetKey() (*ecdsa.PrivateKey, []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
	}
​
	//개인 키를 사용해 공개 키 생성
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
​
	return privateKey, publicKey
}
​
//공개 키 해싱
func HashPubKey(pubKey []byte) []byte {
	//공개 키를 sha256으로 해싱
	pubSHA256 := sha256.Sum256(pubKey)
​
	//sha256으로 해싱한 키를 RIPEND160으로 해싱
	RIPEMDHasher := ripemd160.New()
	_, err := RIPEMDHasher.Write(pubSHA256[:])
	if err != nil {
		fmt.Println(err)
	}
​
	pubRIPEMD160 := RIPEMDHasher.Sum(nil)
	return pubRIPEMD160
}
​
//주소 생성
func GetAddress(pubKey []byte) string {
	pubKeyHash := HashPubKey(pubKey)
	versionPayload := append([]byte{byte(1)}, pubKeyHash...)
	checksum := checksum(versionPayload)
​
	fullload := append(versionPayload, checksum...)
	address := base58.Encode(fullload)
​
	return address
}
​
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:4]
}