package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"math/big"
	//"github.com/syndtr/goleveldb/leveldb/filter"
	//"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	db, err := leveldb.OpenFile("/Users/byunjaejin/Go/level_DB", nil)
	if err != nil {
		panic(err)
	}
	UserId := "aaa"
	hashid := Hash(UserId)
	privKey, pubKey := GetKey()

	//fmt.Printf("개인키 :%x\n", privKey.PublicKey)
	//fmt.Printf("공개키 :%x\n", pubKey)
	println()
	sign := Sign(*privKey, hashid)

	err = db.Put([]byte(UserId), pubKey, nil)

	data, err := db.Get([]byte(UserId), nil)

	bool := Verify(pubKey, sign, hashid)

	fmt.Println("검증 결과 : ", bool)
	fmt.Printf("aaa의 공개키 : %x", data)
	c := Sign(*privKey, []byte{1, 2, 3, 4, 5})
	fmt.Println(Verify(pubKey, c, []byte{1, 2, 3, 4, 5}))
}

// 바이트를 문자열로
/*
func decode(b []byte) string {
	return string(b[:len(b)])
}
*/

//공개 키, 개인 키 생성
func GetKey() (*ecdsa.PrivateKey, []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
	}

	//개인 키를 사용해 공개 키 생성
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return privateKey, publicKey
}

//서명    < 개인키로 해시된 id 암호화 >
func Sign(privKey ecdsa.PrivateKey, hashid []byte) []byte {

	// 개인키로 HashID를 서명한다.
	r, s, err := ecdsa.Sign(rand.Reader, &privKey, hashid)
	if err != nil {
		fmt.Println(err)
	}

	sig := append(r.Bytes(), s.Bytes()...)
	return sig
}

//넘어온 데이터 검증
func Verify(pubKey []byte, Sign []byte, hashid []byte) bool {
	curve := elliptic.P256()

	//서명 데이터 분할
	r := big.Int{}
	s := big.Int{}
	siglen := len(Sign)
	r.SetBytes(Sign[:(siglen / 2)])
	s.SetBytes(Sign[(siglen / 2):])

	//공개키 분할
	x := big.Int{}
	y := big.Int{}
	keylen := len(pubKey)
	x.SetBytes(pubKey[:(keylen / 2)])
	y.SetBytes(pubKey[(keylen / 2):])

	//공개키 찾기
	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	fmt.Printf("rawPubKey : %x\n", rawPubKey)
	//찾은 공개키로 서명 검증
	return ecdsa.Verify(&rawPubKey, hashid, &r, &s)
}

// Hash returns the hash of the Transaction
func Hash(UserID string) []byte {
	var hash [32]byte
	hash = sha256.Sum256([]byte(UserID))
	return hash[:]
}
