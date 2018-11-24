package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)

	for a, b := range s {
		fmt.Println(a, b)
	}
}
