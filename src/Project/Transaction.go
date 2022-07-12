package Project

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	f "fmt"
	"math/big"
	"strconv"
	"time"
)

type Tx struct {
	TxID      []byte `json:"TxID"`      //sha256(Data + TimeStamp + Nonce)
	TimeStamp int64  `json:"TimeStamp"` //Tx 생성시간
	UserID    string `json:"UserID"`    //Tx 발생 시킨 유저 ID
	LogDb     int    `json:"LogDB"`     //LogDB 의 정보
	Content   string `json:"Content"`   //Tx 내용
	RId       int    `json:"RId"`       //Content Type
}

type Txs struct {
	Txs []*Tx `json: Txs`
}

func New_Transaction_Struct() *Txs {
	return &Txs{}
}

func (txs *Txs) AddTx(data []byte) {
	newTx := NewTranscation(data)
	txs.Txs = append(txs.Txs, newTx)
}

//트랜잭션 생성
func NewTranscation(data []byte) *Tx {
	Tx := &Tx{}

	//언마샬하여  Tx 구조체형식으로 저장
	err := json.Unmarshal(data, Tx)
	if err != nil {
		panic(err)
	}

	Tx.TimeStamp = time.Now().Unix()

	timestamp := strconv.FormatInt(Tx.TimeStamp, 10)
	timeBytes := []byte(timestamp)

	//var blockBytes []byte
	blockBytes := append(timeBytes, Tx.UserID...)
	// 		↳--------------↴
	hash := sha256.Sum256(blockBytes)
	Tx.TxID = hash[:]
	//Tx.Sign = hash[:]
	return Tx
}

//트랜잭션 조회
func (txs *Txs) Find_tx(UserID string) *Txs {
	UserTxs := &Txs{}
	for _, v := range txs.Txs {
		if v.UserID == UserID {
			UserTxs.Txs = append(UserTxs.Txs, v)
		}
	}

	return UserTxs

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

	//찾은 공개키로 서명 검증
	return ecdsa.Verify(&rawPubKey, hashid, &r, &s)
}

func (tx *Tx) TPrint(i int) {
	f.Println("-------------------------------", i, "번째 트랜잭션 ---------------------------------")
	// f.Printf("TxID    	 : %x\n", tx.TxID)
	// f.Println("TimeStamp 	 :", time.Unix(tx.TimeStamp, 0))
	// f.Printf("UserID 		 : %s\n", tx.UserID)
	// f.Printf("LogDB		 : %s\n", tx.LogDB)
	// f.Printf("Content 	 : %s\n", tx.Content)
	// f.Printf("RId 	 	 : %s\n", tx.RId)
	// f.Println("-------------------------------------------------------------------------------------\n")
}

func (txs *Txs) Find_Current_tx(Rid int, LogDb int) *Tx {
	CurrentTxs := &Txs{}
	for _, v := range txs.Txs {
		if v.RId == Rid && v.LogDb == LogDb {

			CurrentTxs.Txs = append(CurrentTxs.Txs, v)
		}
	}

	if CurrentTxs.Txs != nil {
		return CurrentTxs.Txs[len(CurrentTxs.Txs)-1]
	} else {
		return nil
	}

}

//저장된 트랜잭션을 참조하고, 요청할 정보가 담긴 구조체를 인자로
func (txs *Txs) GetRTxs(RidTxs *RidTxs) *Txs {
	respTxs := &Txs{}

	for _, v := range RidTxs.Rts {
		for _, v2 := range txs.Txs {
			if v.LogDb == v2.LogDb && v.RId == v2.RId {
				respTxs.Txs = append(respTxs.Txs, v2)
			}
		}
	}

	return respTxs
}
