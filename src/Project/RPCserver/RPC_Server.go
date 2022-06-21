package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	f "fmt"
	base58 "github.com/btcsuite/btcd/btcutil/base58"
	ripemd160 "golang.org/x/crypto/ripemd160"
	"log"
	"net"
	"net/rpc"
	"time"
)

const version = byte(0x00)
const addressChecksumLen = 4

type Calc int //RPC 서버에 등록하기 위한 임의의 타입정의

type Args struct {
	A int
}

type Wallet struct { //받을 값 //지갑
	Timestamp  int64
	Address    string
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

type Wallets struct {
	Wallets map[string]*Wallet
}

var W *Wallets

func init() {
	W = NewWallets()
}

//키 생성 , 지갑 주소 생성
func (c Calc) Get(args Args, reply *Wallet) error {
	private, public := NewKeyPair()
	reply.PrivateKey = private
	reply.PublicKey = public
	reply.Address = reply.GetAddress()
	reply.Timestamp = time.Now().UTC().Unix()

	//지갑 추가
	W.AddWallet(reply.Timestamp, reply.Address, private, public)
	//지갑 추가 확인
	/*
		for _, v := range w.Wallets {
			f.Printf("%d\n", v.Timestamp)
			f.Printf("%s\n", v.Address)
			f.Printf("%x\n", v.PrivateKey)
		}*/
	f.Println(len(W.Wallets))

	return nil
}

func main() {

	gob.Register(elliptic.P256())
	rpc.Register(new(Calc))
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		f.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		f.Println("연결성공")
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

// GetAddress returns wallet address
func (r *Wallet) GetAddress() string {
	pubKeyHash := HashPubKey(r.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := base58.Encode(fullPayload)

	return address
}

// Checksum generates a checksum for a public key
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

// HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func NewWallets() *Wallets {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	return &wallets
}

//지갑 추가
func (ws *Wallets) AddWallet(time int64, addr string, privKey ecdsa.PrivateKey, pubKey []byte) {
	w := &Wallet{time, addr, privKey, pubKey}
	ws.Wallets[w.Address] = w
}
