package main

import "fmt"

func A(pattern string, handler func(int, string))  {
	fmt.Println(pattern)
	handler(1,"a")
}

func B() func(int, string) {

	//fmt.Println(a,b)

	return func(int, string) {

	}
}

func C(func(int, string)){

}

func D(a int, b string){

}

func main() {

	A("test", B())

	C(D)
}
