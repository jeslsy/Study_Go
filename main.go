package main

import (
	"fmt"

	"github.com/jeslsy0507/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(1000)

	err := account.Withdraw(3000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
