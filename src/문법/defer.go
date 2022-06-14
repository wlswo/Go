package main

import (
	f "fmt"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	defer f.Println("defer 1")
	defer f.Println("defer 2")
	defer f.Println("defer 3")
	defer file.Close()

	if err != nil {
		f.Print(err)
	}
	bytes := make([]byte, 100)
	file.Read(bytes)
	f.Print(bytes)
}
