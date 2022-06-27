package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	mux := http.NewServeMux()

	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // 조회 - 하는 함수를 호출(서버에 요청?)
			json.NewEncoder(wr).Encode(logs)
		case http.MethodPost: // 등록 - 하는 함수를 호출(서버에 POST하기)
			var log Log
			json.NewDecoder(r.Body).Decode(&log)

			logs[log.UserId] = &log

			json.NewEncoder(wr).Encode(log)

			fmt.Println(log)

			// 받은 걸 넣어주기. 여기서 포스팅이 들어가면 될것 같음
			pbytes, _ := json.Marshal(log)
			buff := bytes.NewBuffer(pbytes)

			resp, err := http.Post("http://localhost:8081/logs", "application/json", buff)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
		}
	})

	// 3
	mux.Handle("/logs", jsonContentTypeMiddleware(userHandler))
	// mux.HandleFunc("/logs2", func(wr http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("111111")
	// 	fmt.Printf("%v", wr)
	// })
	http.ListenAndServe(":8080", mux)

}
