package Project

import (
	b "bytes"
	//"encoding/json"
	f "fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	UserID  string `json:"UserID"`  //Tx 발생 시킨 유저 ID
	LogDB   string `json:"LogDB"`   //LogDB 의 정보
	Content string `json:"Content"` //Tx 내용
	RId     string `json:"RId"`     //Restaurant 번호
	Sign    []byte `json:"Sign"`    //Id를 Hash + privKey로 암호화한 값
	HashId  []byte `json:"HashId`   //Id를 Hash 한 값
}

func StartTxServer() {
	//서버 키면 트랜잭션 구조체 생성
	txs := New_Transaction_Struct()

	// localhsot:81/create_tx 에 접속시
	http.HandleFunc("/create_tx", func(res http.ResponseWriter, req *http.Request) {

		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			bytes := []byte(respBody)
			txs.CreateTx(bytes) //트랜잭션 생성
		}

		//생성된 트랜잭션의 TxID를 인자값으로 전달 -> 최상위 트랜잭션
		Tx := txs.Txs[len(txs.Txs)-1]

		//Block 생성 서버에 TxID Post 전송
		buff := b.NewBuffer(Tx.TxID)

		resp, err := http.Post("http://localhost:80/newblock", "text/plane", buff)

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		respBody, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			str := string(respBody)
			println(str)
		}

	})

	//http://localhost:81/find_tx 접속시 받는값은 UserID
	http.HandleFunc("/find_tx", func(res http.ResponseWriter, req *http.Request) {

		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			UserID := string(respBody)
			UserTxs := txs.Find_tx(UserID) //트랜잭션 조회
			//구조체 마샬
			//bytes, _ := json.Marshal(UserTxs)
			//buff := b.NewBuffer(bytes)

			for _, v := range UserTxs.Txs {
				f.Printf("TxID : %x\n", v.TxID)
				f.Printf("UserId : %s\n", v.UserID)
			}
			//트랜잭션 조회 결과를 Restful Api로 응답
			// resp, err := http.Post("http://localhost:80/finded_tx", "application/json", buff)

			// if err != nil {
			// 	panic(err)
			// }
			// defer resp.Body.Close()

			// respBody, err = ioutil.ReadAll(resp.Body)
			// if err == nil {
			// 	str := string(respBody)
			// 	println(str)
			// }
		}

		//res.Write([]byte(b)) //웹 브라우저에 응답
	})

	http.ListenAndServe(":81", nil) //80번 포트에서 웹 서버 실행
}
