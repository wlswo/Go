package 문법

//연산자
func Operator() {
	//1. 산술 연산자 +, -, *, /, %
	const c int = 2
	//덧셈 ( + )
	a := 1 + 2
	b := a + c
	str := "Hi, " + " Golang!"

	println(a, b, str)

	//뺄샘 ( - )
	d := a - b
	println(d)

	//곱샘 ( * )
	e := a * b
	f := a / b
	g := b % 2
	println(e, f, g)

	//++, --
	a = 10
	a++
	println(a)

	b = 10
	b--
	println(b)

	c2 := 2 + 2i //복소수 2 + 2i
	c2++
	println(c2)

	//변수 할당과 동시에 ++,-- 연산자 사용 불가능
	/*
		a := 1
		b := a++  //컴파일에러
		c := ++a  //컴파일에러
	*/

}
