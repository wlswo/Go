package main

import "fmt"

func main() {

	//Map 의 key 와 value 를 쌍으로 갖는 자료구조이다.
	// var 변수명 map[key타입]value타입
	/* 맵 선언 종류  */
	var a map[string]int = make(map[string]int)
	//var b = make(map[string]int)
	//b := make(map[int]string)

	a["key1"] = 1
	a["key2"] = 2
	a["key3"] = 3
	//맵 순회
	for i, value := range a {
		fmt.Println(i, value)
	}

	//맵 안에 맵 만들기
	people := map[string]map[string]int{
		"mike": map[string]int{
			"age":    20,
			"height": 170,
		},
		"john": map[string]int{
			"age":    22,
			"height": 175,
		},
	}

	for i, value := range people {
		fmt.Print(i)
		for j, value2 := range value {
			fmt.Println(j, value2)
		}
	}

}
