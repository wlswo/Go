package main

import (
	f "fmt"
	"math/big"
)

func main() {
	i := new(big.Int)
	i.SetString("644", 8) // octal
	f.Println(i)

	b := new(big.Int)
	b.SetString("644", 0)
	f.Println(i.Cmp(b))
}
