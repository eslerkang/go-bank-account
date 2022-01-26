package main

import (
	"fmt"

	"github.com/eslerkang/learngo/accounts"
)




func main() {
	account := accounts.NewAccount("kang")
	account.Deposit(100)
	err:=account.WithDraw(110)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(account.Owner(), account.Balance())
	fmt.Println(account)
}