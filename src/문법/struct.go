package main

import "fmt"

func (r *Rectangle) rect_Calc() int {
	result := r.width * r.height
	return result
}

//구조체를 main 밖에 선언하냐 안에 선언하냐에 따라
//참조할수 있는 범위가 바뀜 지역 or 전역
type Rectangle struct {
	width  int
	height int
}

func main() {

	myRect := Rectangle{30, 40}

	fmt.Println(myRect.rect_Calc())

}
