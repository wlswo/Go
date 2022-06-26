package Project

import (
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

func StartServer() {
	//서버 키면 블록체인 구조체, 트랜잭션 구조체 생성
	bc := NewBlockchain()
	txs := New_Transaction_Struct()

	// localhsot:80/test 에 접속시
	http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		/* 처리할 기능들 작성 */
		//Sample Data
		data := Data{"aaa", "walletDB", "A가 B를 수정", "A타입 게시글"}

		//마샬링
		bytes, _ := json.Marshal(data)

		txs.TxRun(bytes) //트랜잭션 생성
		//생성된 트랜잭션의 TxID를 인자값으로 전달
		Tx := txs.Txs[len(txs.Txs)-1]
		bc.Run(Tx.TxID) //블록 생성

		b, _ := ioutil.ReadFile("blockFile.json") //file read
		/* ------------- */
		res.Write([]byte(b)) //웹 브라우저에 응답
	})

	http.ListenAndServe(":80", nil) //80번 포트에서 웹 서버 실행
}
