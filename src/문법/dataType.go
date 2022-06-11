package 문법

func DataType() {
	//작은 따옴표안에서는 escape문자 그대로 인식하여 출력된다.
	rawLiteral := `아리랑\n 아라리요
				   아리랑\n`
	println(rawLiteral)

	//큰 따옴표 안에서는 escape 문자를 인식한다.
	interLiteral := "아리랑 아리랑\n 아라리요"
	println(interLiteral)

	/* 데이터 타입 변환 (Type Conversion)
	* Type(value) 와 같이 표현한다
	* Go언어에서는 형 변환을 할 때 변환을 명시적으로 지정해주어야한다.
	 */
	//정수 100을 float 타입으로 변경하기
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	//문자열을 바이트배열로 변경하기
	str := "ABC"
	bytes := []byte(str)  //바이트배열로 변경
	str2 := string(bytes) //다시 문자열로 변경
	println(bytes, str2+"\n")
}
