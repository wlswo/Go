package main

import (
	"Project"
	"fmt"
)

func main() {

	wallet := Project.NewWallet()
	address := wallet.GetAddress()

	fmt.Println(Project.ValidateAddress(address))
	fmt.Println(address)

}
