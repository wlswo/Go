package main

import (
	b "bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Id_Key struct {
	UserID string `json:"UserID"` //Tx 발생 시킨 유저 ID
	PubKey []byte `json:"PubKey"` //PubKey
}

func main() {
	_, pubKey := GetKey()

	data := Id_Key{"aaa", pubKey}
	//마샬링
	bytes, _ := json.Marshal(data)
	buff := b.NewBuffer([]byte(bytes))
	// 1. http://localhost:80/save_key 주소로 요청
	// 2. application/json 포맷으로
	// 3. buff 데이터를

	resp, err := http.Post("http://localhost:80/save_key", "application/json", buff)

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
