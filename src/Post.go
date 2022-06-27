package main

import (
	b "bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	UserID  string `json:"UserID"`  //Tx 발생 시킨 유저 ID
	LogDB   string `json:"LogDB"`   //LogDB 의 정보
	Content string `json:"Content"` //Tx 내용
	Ctype   string `json:"Ctype"`   //Content Type
}

func main() {

	data := Data{"aaa", "walletDB", "A가 B를 수정", "A타입 게시글"}

	//마샬링
	bytes, _ := json.Marshal(data)
	buff := b.NewBuffer(bytes)
	// 1. http://localhost:81/create_tx 주소로 요청
	// 2. application/json 포맷으로
	// 3. buff 데이터를

	resp, err := http.Post("http://localhost:81/create_tx", "application/json", buff)

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
	//Project.StartServer()

}
