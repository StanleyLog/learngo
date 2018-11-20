package main

import "fmt"

func main() {
	arr1 := [...]int {1,2,3,4,5}
	slice1 := arr1[:]
	fmt.Println(arr1)
	fmt.Println(slice1)
	update(slice1)
	fmt.Println(arr1)
	fmt.Println(slice1)

	slice2 := []string {"a", "b", "c", "d", "e"}
	fmt.Println(slice2)
}

func update(s []int)  {
	s[0] =100
}
