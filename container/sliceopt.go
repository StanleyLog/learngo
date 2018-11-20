package main

import "fmt"

func main() {

	original_slice := make([]int , 10, 32)
	slice1 := []int {1,2,3}

	fmt.Println(original_slice)
	copy(original_slice, slice1)
	fmt.Println(original_slice)

}
