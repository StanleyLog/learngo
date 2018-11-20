package main

import "fmt"

func main() {

	score := 80

	switch  {
	case score < 60:
		fmt.Println("F")
	case score < 80:
		fmt.Println("C")
	case score < 90:
		fmt.Println("B")
	case score <= 100:
		fmt.Println("A")
	default:
		panic(fmt.Sprintf("score is error! score:%s\\n ", score))
	}


}
