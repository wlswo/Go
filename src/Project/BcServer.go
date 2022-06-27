package Project

import (
	//b "bytes"
	//	"encoding/json"
	//f "fmt"
	"io/ioutil"
	"net/http"
)

func StartBCServer() {
	//서버 키면 블록체인 구조체 생성
	bc := NewBlockchain()

	// localhsot:80/test 에 접속시
	http.HandleFunc("/create_bc", func(res http.ResponseWriter, req *http.Request) {
		respBody, err := ioutil.ReadAll(req.Body)
		if err == nil {
			TxID := []byte(respBody)
			/* 처리할 기능들 작성 */
			bc.CreateBc(TxID) //블록 생성
		}
		defer req.Body.Close()

		//b, _ := ioutil.ReadFile("BlockFile.json") //file read
		//res.Write([]byte(b)) //웹 브라우저에 응답
		// /*------ PBFT 서버에 POST 전달 -----*/
		// Block := bc.Blocks[len(bc.Blocks)-1]
		// bytes, _ := json.Marshal(Block)
		// buff := b.NewBuffer(bytes)

		// resp, err := http.Post("http://localhost:4000/pbft", "application/json", buff)

		// if err != nil {
		// 	panic(err)
		// }

		// defer resp.Body.Close()

		// respBody, err = ioutil.ReadAll(resp.Body)
		// if err == nil {
		// 	str := string(respBody)
		// 	println(str)
		// }
		// /*---------------*/
	})

	// //pbft -> Here Responce
	// http.HandleFunc("/reply", func(res http.ResponseWriter, req *http.Request) {
	// 	respBody, err := ioutil.ReadAll(req.Body)
	// 	if err == nil {
	// 		result := []byte(respBody)
	// 		f.Println(result)
	// 	}
	// 	defer req.Body.Close()
	// })

	http.ListenAndServe(":80", nil) //80번 포트에서 웹 서버 실행
}
