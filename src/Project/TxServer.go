package Project

import (
	b "bytes"
	"io/ioutil"
	"net/http"
)

type Data struct {
	UserID  string `json:"UserID"`  //Tx 발생 시킨 유저 ID
	LogDB   string `json:"LogDB"`   //LogDB 의 정보
	Content string `json:"Content"` //Tx 내용
	Ctype   string `json:"Ctype"`   //Content Type
}

func StartTxServer() {
	//서버 키면 트랜잭션 구조체 생성
	txs := New_Transaction_Struct()

	// localhsot:80/test 에 접속시
	http.HandleFunc("/create_tx", func(res http.ResponseWriter, req *http.Request) {

		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			bytes := []byte(respBody)
			txs.CreateTx(bytes) //트랜잭션 생성
		}

		//생성된 트랜잭션의 TxID를 인자값으로 전달
		Tx := txs.Txs[len(txs.Txs)-1]
		// //Tx.TxID 만 넘기기
		//TxFile, _ := ioutil.ReadFile("TxFile.json") //file read

		//Block 생성 서버에 TxID Post 전송
		buff := b.NewBuffer(Tx.TxID)

		resp, err := http.Post("http://localhost:80/create_bc", "text/plane", buff)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		respBody, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			str := string(respBody)
			println(str)
		}

		//res.Write([]byte(b)) //웹 브라우저에 응답
	})

	http.ListenAndServe(":81", nil) //80번 포트에서 웹 서버 실행
}
