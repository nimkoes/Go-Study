package main

import (
	"accounts"
	"fmt"
	"log"
)

func main() {
	myAccount := accounts.NewAccount("Nimkoes")
	myAccount.Deposit(10)
	fmt.Println(myAccount.Balance())

	err := myAccount.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(myAccount.Balance())
}
