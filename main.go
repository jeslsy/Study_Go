package main

import (
	"fmt"

	"github.com/jeslsy0507/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(1000)
	fmt.Println(account)
}
