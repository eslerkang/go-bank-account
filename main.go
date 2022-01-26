package main

import (
	"fmt"
	"log"

	"github.com/eslerkang/learngo/accounts"
)




func main() {
	account := accounts.NewAccount("kang")
	account.Deposit(100)
	fmt.Println(*account)
	err:=account.WithDraw(110)
	if err!=nil {
		log.Fatal(err)
	}
}