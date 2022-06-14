package main

import (
	"fmt"
	"strings"
)

//Input: String
//Output: []byte
//
// len(In) = 13* 4 + 2 = 54
// In: f03963DefABcD42b55060a6f688
//     025b7de0b31777614ce174b8356
//     51843e301b64a52212d3226adc2
//     3a4545f1b204358bda427530920
//
// len(Out) = 27
// Out: [0]byte = 0xf0
//      [1]  = 0x39
//      [2]  = 0x63
//
//

func main() {
	ciphertext := toByte("f03963DefABcD42b55060a6f688025b7de0b31777614ce174b835651843e301b64a52212d3226adc23a4545f1b204358bda427530920")
	fmt.Println(ciphertext)
	fmt.Printf("0x%x ", ciphertext)
}

func toByte(s string) []byte {
	//1. 입력받은 string 값을 2개 단위로 끊어서 저장 받기
	//2. 입력받은 2개의값을 0x + s 상태에서 byte 값으로 변환

	//소문자로 변환과 byte 배열로
	bs := []byte(strings.ToLower(s))
	b := make([]byte, len(bs)/2)

	//배열 순회 , 2개씩받아주면서
	for i := 0; i < len(s)/2; i++ {
		switch bs[i*2] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b[i] = (bs[i*2] - '0') << 4
		case 'a', 'b', 'c', 'd', 'e', 'f':
			b[i] = (bs[i*2] - 'a' + 10) << 4
		}
		switch bs[i*2+1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b[i] += (bs[i*2] - '0') << 4
		case 'a', 'b', 'c', 'd', 'e', 'f':
			b[i] += (bs[i*2] - 'a' + 10) << 4
		}
	}

	return b
}
