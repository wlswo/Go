package main

import (
	b "bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"time"
)

type Data struct {
	UserID  string `json:"UserID"`  //Tx 발생 시킨 유저 ID
	LogDB   string `json:"LogDB"`   //LogDB 의 정보
	Content string `json:"Content"` //Tx 내용
	RId     string `json:"RId"`     //Content Type
	Sign    []byte `json:"Sign"`    //Signature <= UserID 를 SHA256() 해시화 하여 개인키로 (ecdsa.Sign() 함수 이용 ) 암호화한 값
	HashId  []byte `json:"HashId"`  //Id를 sha256 돌린값
}

func main() {
	for i := 0; i < 3000; i++ {
		//0.7 sec
		//time.Sleep(time.Nanosecond * 500000000)
		//time.Sleep(time.Second * 1)
		privKey, _ := GetKey()
		//Id 해시
		hashId := Hash("aaa")
		//해시된 Id 공개키로 암호화
		r, s, err := ecdsa.Sign(rand.Reader, privKey, hashId[:])
		//r,s 더한값이 서명
		Signature := append(r.Bytes(), s.Bytes()...)

		data := Data{"aaa", "walletDB", "A가 B를 수정2", "Restorant ID", Signature, hashId}

		//마샬링
		bytes, _ := json.Marshal(data)
		buff := b.NewBuffer([]byte(bytes))
		// 1. http://localhost:81/create_tx 주소로 요청
		// 2. application/json 포맷으로
		// 3. buff 데이터를

		resp, err := http.Post("http://localhost:80/create_bc", "application/json", buff)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		// Response 체크.
		respBody, err2 := ioutil.ReadAll(resp.Body)
		if err2 == nil {
			str := string(respBody)
			println(str)
		}
	}

}

//공개 키, 개인 키 생성
func GetKey() (*ecdsa.PrivateKey, []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		println(err)
	}

	//개인 키를 사용해 공개 키 생성
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return privateKey, publicKey
}

// Hash returns the hash of the Transaction
func Hash(UserID string) []byte {
	var hash [32]byte
	hash = sha256.Sum256([]byte(UserID))
	return hash[:]
}
