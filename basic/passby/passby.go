package main

import "fmt"

func pass_by(a int) int{
	a++

	return a
}

func main() {

	var a int = 0

	fmt.Println(a)
	fmt.Println(pass_by(a))
	fmt.Println(a)
}
