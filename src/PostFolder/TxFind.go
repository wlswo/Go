package main

import (
	b "bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"time"
)

type CurrentTxData struct {
	RId   int `json:"RId"`   //Restaurant 번호
	LogDb int `json:"LogDB"` //LogDB 의 정보
}

func main() {
	data := CurrentTxData{1, 2}
	bytes, _ := json.Marshal(data)
	buff := b.NewBuffer([]byte(bytes))

	resp, err := http.Post("http://localhost:81/current_tx", "text/plane", buff)
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
