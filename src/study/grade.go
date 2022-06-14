package main

import f "fmt"

func main() {
	//scan	 공백과 개행으로 구분하여 입력
	//scnaln 공백으로 구분하여 입력
	//scanf  포멧 지정자를 이용하여 개발자가 원하는 형태로 입력가능

	var score int
	f.Scan(&score)

	switch {
	case score > 90 && score < 100:
		f.Println("A")
	case score >= 80:
		f.Println("B")
	case score >= 70:
		f.Println("C")
	default:
		f.Println("D")

	}

}
