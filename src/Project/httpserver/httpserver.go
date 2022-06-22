package main

import (
	P "Project"
	"io/ioutil"
	"net/http"
)

func main() {
	//서버 키면 블록체인 구조체 생성
	bc := P.NewBlockchain()
	// localhsot:80/test 에 접속시
	http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		/* 처리할 기능들 작성 */

		bc.Run("good") //블록 생성

		b, _ := ioutil.ReadFile("blockFile.json") //file read
		/* ------------- */
		res.Write([]byte(b)) //웹 브라우저에 응답
	})

	http.ListenAndServe(":80", nil) //80번 포트에서 웹 서버 실행
}
