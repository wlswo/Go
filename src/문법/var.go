package 문법

func Var() {
	/*
	*  변수는 Go 키워드 var를 사용
	*  var [변수명] [변수타입]
	 */

	var a int = 11

	var f float32 = 11 //초기값 선언

	a = 10
	f = 12 //변수값 재 선언

	println(a, f)

	//복수 선언 가능
	var i, j, k int = 1, 2, 3
	print(i, j, k)

	/* 초기값을 지정하지 않을경우 "Zero Value"를 기본적으로 할당함
	 * 숫자형은 0
	 * bool 타입은 false
	 * String 형에는 "" (빈문자열)
	 */

	/* 변수타입을 자동으로 추론해주는 기능 존재
	 * i 에는 정수형으로 s 에는 문자열로 할당됨
	 */

	var i2 = 1
	var s2 = "HI"
	print(i2, s2)
	/* 변수를 선언하는 또 다른 방식으론
	 * Short Assignment Statement ( := )를 사용할수 있음.
	 * 즉, var i = 1 대신 i := 1 로 var를 생략가능
	 * 하지만 이러한 표현은 함수(func) 내에서만 사용할 수 있으며
	 * 함수 밖에서는 var를 사용해야 한다.
	 * Go에서 변수와 상수는 함수 밖에서도 사용할 수 있다.
	 */

	// i := 1

	/* 상수
	* const [변수명] [변수타입]
	 */

	const c int = 10
	const s string = "Hi"

	//상수도 변수타입 추론 기능이 존재한다.
	const c3 = 10
	const s3 = "Hi"

	// 여러개의 상수들을 묶어서 지정가능
	const (
		name   = "GilDong"
		age    = "27"
		gender = "Man"
	)

	//꿀팁 상수값을 0 부터 순차적으로 부여하는 iota라는 identifier가 있다.
	const (
		Apple  = iota //0
		Grape         //1
		Orange        //2
	)
}
