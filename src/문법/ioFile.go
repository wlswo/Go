package main

import (
	f "fmt"
	"os"
)

func main() {
	/* File Create , Wrtie */
	file1, _ := os.Create("hello1.txt")
	defer file1.Close()

	f.Fprint(file1, 1, 1.1, "Hello,World1!")

	file2, _ := os.Create("hello2.txt")
	defer file2.Close()
	f.Fprintf(file2, "%d,%f,%s", 1, 1.1, "Hello, world2!") //File print Format => 파일에 형식 대로 저장

	/* File Read, Input */
	var num1 int
	var num2 float32
	var s string

	Read_file1, _ := os.Open("hello1.txt")
	n, _ := f.Fscan(Read_file1, &num1, &num2, &s)
	f.Println("입력받은개수 : ", n)
	f.Println(num1, num2, s)

}
