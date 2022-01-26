package main

import (
	"fmt"

	"github.com/eslerkang/learngo/dict"
)




func main() {
	myDict := dict.Dictionary{"first": "first word"}
	word:="hello"
	definition:="word"
	err:=myDict.Add(word, definition)
	if err!=nil {
		fmt.Println(err)
	}
	def, err := myDict.Search(word)
	fmt.Println(def)

	err = myDict.Add(word, definition)
	if err!=nil {
		fmt.Println(err)
	}
}