package main

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	f "fmt"
	"math/big"
	mrand "math/rand"
	"strconv"
	"time"
)

func main() {

	//로또 번호
	//무작위 번호 6개 출력 포인트는 중복 제거
	/* crypto/rand 함수 이용 */
	for i := 0; i < 6; i++ {
		random, error := crand.Int(crand.Reader, big.NewInt(45))
		if true {
			f.Print(random, " ")
		} else {
			f.Print(error)
		}
	}
	f.Println("")
	/* math/rand 함수 이용 */
	mrand.Seed(time.Now().UnixNano()) //가변값을 가진 시드를 주어야 math/rand 값이 바뀜

	switch i := mrand.Intn(10); {
	case i >= 3 && i < 6:
		f.Println("i의 값은 3 이상 6미만")
	case i == 9:
		f.Println("i의 값은 9")
	default:
		f.Println(i)
	}
	f.Println()
	/* 랜덤값 받아서 해시 변환 해보기 */
	hash := sha256.New() //해쉬 함수선언
	number := mrand.Intn(100)
	data := strconv.Itoa(number)

	hash.Write([]byte(data))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	f.Println("-------- 랜덤 값을 받아 해쉬 변환 -------")
	f.Println(mdStr, "\n")

}
