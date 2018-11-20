package main

import "learngo/functional/fib"

func main() {
	a := fib.Fibonacci()

	//for i := 0; i < 10; i++ {
	//	fmt.Println(a())
	//}
	//
	//b := fmt.Sprintf("ffffaaaa%d, %s", 1111, "sdfsdfsdf")
	//fmt.Println(b)

	fib.PrintFileContents(a)

}
