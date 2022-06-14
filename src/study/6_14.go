package main

import f "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6}

	//향상된 포문 같은 for 문
	for i, value := range a {
		f.Println("a[", i, "]", ":", value)
	}
}
