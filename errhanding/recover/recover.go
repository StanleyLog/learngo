package main

import "fmt"

func tryRecover()  {

	defer func() {

		r := recover()

		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}

	}()

	b := 0
	a := 1/b
	fmt.Println(a)

	//panic("12312")
}

func main() {

	tryRecover()

	fmt.Println("end of process")

}


