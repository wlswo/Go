package main

import f "fmt"

func main() {

	//익명함수 변수를 함수처럼 사용가능
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	sub := func(a int, b int) int {
		return a - b
	}

	/* 일급함수 함수의 파라미터로 함수를 전달시킬수가 있다.
	*  그렇다면 어떻게 전달해야 할까
	*  익명함수를 이용하여 변수에 함수를 저장시키면 쉽게 전달시킬수 있다는 것이다.
	*  리턴값 또한 마찬가지 이다.
	 */
	result := cal(sum, 10)
	f.Println(result)
	result = sub_Calc(sub, 10, 5)
	f.Println(result)
}

func cal(f func(n ...int) int, x int) int {
	result := f(1, 2, 3, x)

	return result
}

func sub_Calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}
