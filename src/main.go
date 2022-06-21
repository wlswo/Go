package main

import (
	"Project"
	"fmt"
)

func main() {
	Project.Run()

	wallet := Project.NewWallet()
	address := wallet.GetAddress()

	fmt.Println(Project.ValidateAddress(address))
	fmt.Println(address)
	/* 트랜잭션들 관리 */
	/* 트랜잭션에 대한 서명 */

}
