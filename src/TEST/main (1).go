package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var logs = map[string]*Log{}

type Log struct {
	UserId  string `json:"UserId"`
	LogDb   string `json:"LogDb"`
	RId     int    `json:"RId"`
	Content string `json:"Content"`
}

// 4
func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	// 들어오는 요청의 Response Header에 Content-Type을 추가
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		// 전달 받은 http.Handler를 호출한다.
		next.ServeHTTP(rw, r)
		// fmt.Printf("%v", rw)
		// http.Post("localhost:8080/logs2", "application/json", r.Body)
	})
}

func main() {

	// 텍스트로 포스팅 해보기
	// log := "{\"UserId\": \"aaa\",\"LogDb\": \"log_r_name\",\"RId\": 5,\"Content\": \"버거킹\"}"

	// reqBody := bytes.NewBufferString(log)

	// resp, err := http.Post("http://127.0.0.1:8080/logs", "text/plain", reqBody)

	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	mux := http.NewServeMux()

	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		switch r.Method {
		case http.MethodGet: // 조회 - 하는 함수를 호출(서버에 요청?)
			var log Log
			err := json.Unmarshal(b, &log)
			if err != nil {
				panic(err)
			}
			fmt.Println(log)

		case http.MethodPost:
			var log Log
			err := json.Unmarshal(b, &log)
			if err != nil {
				panic(err)
			}
			fmt.Println(log)
		}
	})

	// 3
	mux.Handle("/logs", jsonContentTypeMiddleware(userHandler))
	// mux.HandleFunc("/logs2", func(wr http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("111111")
	// 	fmt.Printf("%v", wr)
	// })
	http.ListenAndServe(":8081", mux)
}
