package main

import (
	"net/http"
)

func main() {
	s := "Hello World!"

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		//HTML웹 페이지 작성
		html := `
		<html>
		<head>
			<title>Hello</title>
			<script type="text/javascript" src="/assets/hello.js"></script>
			<link href="/assets/hello.css" rel="stylesheet" />
		</head>
		<body>
			<span class="hello">` + s + `</span>
		</body>
		</html>
		`
		res.Header().Set("Content-Type", "text/html") //HTML 헤더 설정
		res.Write([]byte(html))                       //웹 브라우저에 응답
	})

	http.Handle( // /assets/ 경로에 도착시 파일 서버를 동작시킴
		"/assets/",
		http.StripPrefix( //파일 서버를 실행할 때 assets
			//디렉터리를 지정했으므로 URL 경로에서 /assets/ 삭제
			"/assets/",
			http.FileServer(http.Dir("assets")), //웹 서버에서 assets 디렉터리
		),
	)
	http.ListenAndServe(":80", nil) //80번 포트에서 웹 서버 실행
}
