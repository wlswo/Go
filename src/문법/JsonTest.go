package main

/*
	marshalling : 결집시키다 , 모으다

	파일을 데이터나 , 바이트 스트림 즉 바이너리 형태로 변환 하는 작업을 일컫
*/

import (
	"encoding/json"
	"fmt"
)

func main() {
	//json
	txt := `{
		"name" : "홍길동",
		"age"  : 10
	}
	`
	//		   [Key Type][Value Type]
	var data map[string]interface{} //<--JSON 문서의 데이터를 저장할 공간
	//   ㄴ---------------------ㄱ
	// 		        			↴
	json.Unmarshal([]byte(txt), &data)

	fmt.Println(data["name"], data["age"])
}
