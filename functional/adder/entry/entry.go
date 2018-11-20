package main

import (
	"fmt"
	"learngo/functional/adder"
)

func main() {

	//var a int = 1
	//b := test(a)

	f :=  adder.Adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %d, v:%d\n", i, f(i))
	}

}